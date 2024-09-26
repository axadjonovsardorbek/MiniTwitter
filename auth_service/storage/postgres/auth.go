package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/axadjonovsardorbek/MiniTwitter/auth_service/api/token"
	pb "github.com/axadjonovsardorbek/MiniTwitter/auth_service/genproto/auth"
	"github.com/google/uuid"

	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepo struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewAuthRepo(db *sql.DB, rdb *redis.Client) *AuthRepo {
	return &AuthRepo{db: db, rdb: rdb}
}

func (r *AuthRepo) Create(req *pb.UserCreateReq) (*pb.TokenRes, error) {

	id := uuid.New().String()

	query := `
	INSERT INTO users(
		id,
		name,
		username,
		bio,
		image_url,
		password,
		email
	) VALUES($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.Exec(query, id, req.Name, req.Username, req.Bio, req.ImageUrl, req.Password, req.Email)

	if err != nil {
		fmt.Println("error while registering")
		return nil, errors.New("error while registering")
	}

	tkn := token.GenerateJWTToken(id, "user", req.Email)

	return &pb.TokenRes{
		AccessToken:  tkn.AccessToken,
		RefreshToken: tkn.RefreshToken,
		Id:           id,
	}, nil
}
func (r *AuthRepo) GetById(req *pb.ById) (*pb.UserRes, error) {

	tz, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			tz.Rollback()
			err = fmt.Errorf("panic occurred: %v", p)
		} else if err != nil {
			tz.Rollback()
		} else {
			err = tz.Commit()
		}
	}()

	user := pb.UserRes{}

	query := `
	SELECT
		id,
		name,
		username,
		CASE 
			WHEN bio IS NULL THEN '' 
			ELSE bio::text 
		END as default_bio,	
		CASE 
			WHEN image_url IS NULL THEN '' 
			ELSE image_url::text 
		END as default_image_url,	
		email,
		to_char(created_at, 'YYYY-MM-DD')
	FROM 
		users
	WHERE
		id = $1
	AND 
		deleted_at = 0
	`

	row := tz.QueryRow(query, req.Id)

	err = row.Scan(
		&user.Id,
		&user.Name,
		&user.Username,
		&user.Bio,
		&user.ImageUrl,
		&user.Email,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	queryFollow := `
	SELECT 
		COUNT(follower_id)
	FROM
		follows
	WHERE
		follower_id = $1
	AND	
		deleted_at = 0
	`

	row = tz.QueryRow(queryFollow, req.Id)

	err = row.Scan(
		&user.Follows,
	)

	if err != nil {
		return nil, err
	}

	queryFollower := `
	SELECT 
		COUNT(followed_id)
	FROM
		follows
	WHERE
		followed_id = $1
	AND	
		deleted_at = 0
	`

	row = tz.QueryRow(queryFollower, req.Id)

	err = row.Scan(
		&user.Followers,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (r *AuthRepo) Update(req *pb.UserUpdateReq) (*pb.Void, error) {

	query := `
	UPDATE 
		users
	SET
	`

	var conditions []string
	var args []interface{}

	if req.Name != "" && req.Name != "string" {
		conditions = append(conditions, " name = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Name)
	}
	if req.Username != "" && req.Username != "string" {
		conditions = append(conditions, " username = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Username)
	}
	if req.Bio != "" && req.Bio != "string" {
		conditions = append(conditions, " bio = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Bio)
	}
	if req.ImageUrl != "" && req.ImageUrl != "string" {
		conditions = append(conditions, " image_url = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.ImageUrl)
	}
	if req.Email != "" && req.Email != "string" {
		conditions = append(conditions, " email = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Email)
	}

	if len(conditions) == 0 {
		return nil, errors.New("nothing to update")
	}

	conditions = append(conditions, " updated_at = now()")
	query += strings.Join(conditions, ", ")
	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0"

	args = append(args, req.Id)

	res, err := r.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("user with id %s couldn't update", req.Id)
	}

	return &pb.Void{}, nil
}
func (r *AuthRepo) Delete(req *pb.ById) (*pb.Void, error) {
	query := `
	UPDATE
		users
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := r.db.Exec(query, req.Id)

	if err != nil {
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("user with id %s not found", req.Id)
	}

	return &pb.Void{}, nil
}
func (r *AuthRepo) ChangePassword(req *pb.ChangePasswordReq) (*pb.Void, error) {
	var curPass string

	queryCurrent := `SELECT password FROM users WHERE id = $1 AND deleted_at = 0`

	row := r.db.QueryRow(queryCurrent, req.Id)

	err := row.Scan(&curPass)

	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(curPass), []byte(req.OldPassword)); err != nil {
		return nil, errors.New("invalid current password")
	}

	queryUpdate := `UPDATE users SET password = $1, updated_at = now() WHERE id = $2 AND deleted_at = 0`

	_, err = r.db.Exec(queryUpdate, req.NewPassword, req.Id)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
func (r *AuthRepo) Login(req *pb.LoginReq) (*pb.TokenRes, error) {
	var id string
	var username string
	var password string
	var email string

	query := `
	SELECT 
		id,
		username,
		password,
		email
	FROM
		users
	WHERE
		deleted_at = 0
	AND 
		username = $1
	`

	row := r.db.QueryRow(query, req.Username)

	err := row.Scan(
		&id,
		&username,
		&password,
		&email,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	if err != nil {
		log.Println("Error while login user: ", err)
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	tkn := token.GenerateJWTToken(id, "user", email)

	return &pb.TokenRes{
		AccessToken:  tkn.AccessToken,
		RefreshToken: tkn.RefreshToken,
		Id:           id,
	}, nil
}
