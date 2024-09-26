package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	pb "github.com/axadjonovsardorbek/MiniTwitter/twit_service/genproto/twit"
	"github.com/google/uuid"
)

type TwitRepo struct {
	db *sql.DB
}

func NewTwitRepo(db *sql.DB) *TwitRepo {
	return &TwitRepo{db: db}
}

func (r *TwitRepo) Create(req *pb.TwitCreateReq) (*pb.Void, error) {
	id := uuid.New().String()

	query := `
	INSERT INTO twits(
		id,
		user_id,
		twit_id,
		content,
		image_url
	) VALUES($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(query, id, req.UserId, req.TwitId, req.Content, req.ImageUrl)

	if err != nil {
		fmt.Println("error while creating twit")
		return nil, errors.New("error while creating twit")
	}

	return nil, nil
}
func (r *TwitRepo) GetById(req *pb.ById) (*pb.TwitRes, error) {

	res := pb.TwitRes{}

	query := `
	SELECT
		id,
		user_id,
		twit_id,
		content,
		image_url
	FROM
		twits
	WHERE 
		deleted_at = 0
	AND
		user_id = $1
	`

	row := r.db.QueryRow(query, req.Id)
	err := row.Scan(
		&res.Id,
		&res.UserId,
		&res.TwitId,
		&res.Content,
		&res.ImageUrl,
	)
	if err != nil {
		return nil, err
	}

	queryFollow := `
	SELECT 
		COUNT(twit_id)
	FROM
		likes
	WHERE
		twit_id = $1
	AND	
		deleted_at = 0
	`

	row = r.db.QueryRow(queryFollow, req.Id)

	err = row.Scan(
		&res.Likes,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
func (r *TwitRepo) GetAll(req *pb.TwitGetAllReq) (*pb.TwitGetAllRes, error) {
	res := &pb.TwitGetAllRes{}

	query := `
	SELECT
		COUNT(id) OVER () AS total_count,
		id,
		user_id,
		twit_id,
		content,
		image_url
	FROM
		likes
	WHERE 
		deleted_at = 0
	`

	var args []interface{}
	var conditions []string

	if req.UserId != "" && req.UserId != "string" {
		conditions = append(conditions, fmt.Sprintf("user_id = $%d", len(args)+1))
		args = append(args, req.UserId)
	}

	if req.TwitId != "" && req.TwitId != "string" {
		conditions = append(conditions, fmt.Sprintf("twit_id = $%d", len(args)+1))
		args = append(args, req.TwitId)
	}
	if req.Content != "" && req.Content != "string" {
		conditions = append(conditions, fmt.Sprintf("LOWER(content) LIKE LOWER($%d", len(args)+1)+")")
		args = append(args, "%"+req.Content+"%")
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("twits list is empty")
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		twit := pb.TwitRes{}
		var count int32

		err := rows.Scan(
			&count,
			&twit.Id,
			&twit.UserId,
			&twit.TwitId,
			&twit.Content,
			&twit.ImageUrl,
		)
		if err != nil {
			return nil, err
		}

		queryFollow := `
		SELECT 
			COUNT(twit_id)
		FROM
			likes
		WHERE
			twit_id = $1
		AND	
			deleted_at = 0
		`
	
		row := r.db.QueryRow(queryFollow, twit.Id)
	
		err = row.Scan(
			&twit.Likes,
		)
	
		if err != nil {
			return nil, err
		}

		res.Twits = append(res.Twits, &twit)
	}

	return res, nil
}
func (r *TwitRepo) Update(req *pb.TwitUpdateReq) (*pb.Void, error) {
	res := &pb.Void{}

	query := `
	UPDATE
		twits
	SET`

	var conditions []string
	var args []interface{}

	if req.Content != "" && req.Content != "string" {
		conditions = append(conditions, " content = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Content)
	}
	if req.ImageUrl != "" && req.ImageUrl != "string" {
		conditions = append(conditions, " image_url = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.ImageUrl)
	}

	if len(conditions) == 0 {
		return nil, errors.New("nothing to update")
	}

	conditions = append(conditions, " updated_at = now()")
	query += strings.Join(conditions, ", ")
	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0"
	query += " AND user_id = $" + strconv.Itoa(len(args)+1)

	args = append(args, req.Id)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (r *TwitRepo) Delete(req *pb.TwitDeleteReq) (*pb.Void, error) {

	query := `
	UPDATE 
		twits
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		user_id = $1
	AND 
		id = $2
	AND
		deleted_at = 0
	`

	res, err := r.db.Exec(query, req.UserId, req.Id)

	if err != nil {
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("twit with user_id %s and twit_id %s not found", req.UserId, req.Id)
	}

	return &pb.Void{}, nil
}
