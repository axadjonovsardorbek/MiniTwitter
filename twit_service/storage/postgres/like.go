package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	pb "github.com/axadjonovsardorbek/MiniTwitter/twit_service/genproto/twit"
	"github.com/google/uuid"
)

type LikeRepo struct {
	db *sql.DB
}

func NewLikeRepo(db *sql.DB) *LikeRepo {
	return &LikeRepo{db: db}
}

func (r *LikeRepo) Create(req *pb.LikeCreateReq) (*pb.Void, error) {
	id := uuid.New().String()

	query := `
	INSERT INTO likes(
		id,
		user_id,
		twit_id
	) VALUES($1, $2, $3)
	`

	_, err := r.db.Exec(query, id, req.UserId, req.TwitId)

	if err != nil {
		fmt.Println("error while liking")
		return nil, errors.New("error while liking")
	}

	return nil, nil
}
func (r *LikeRepo) GetAll(req *pb.LikeGetAllReq) (*pb.LikeGetAllRes, error) {
	res := &pb.LikeGetAllRes{}

	query := `
	SELECT
		COUNT(id) OVER () AS total_count,
		id,
		user_id,
		twit_id
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

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("likes list is empty")
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		like := pb.LikeRes{}
		var count int32

		err := rows.Scan(
			&count,
			&like.Id,
			&like.UserId,
			&like.TwitId,
		)
		if err != nil {
			return nil, err
		}

		res.Likes = append(res.Likes, &like)
		res.Count = count
	}

	return res, nil
}
func (r *LikeRepo) Delete(req *pb.LikeCreateReq) (*pb.Void, error) {

	query := `
	UPDATE 
		likes
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		user_id = $1
	AND 
		twit_id = $2
	AND
		deleted_at = 0
	`

	res, err := r.db.Exec(query, req.UserId, req.TwitId)

	if err != nil {
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("like with user_id %s and twit_id %s not found", req.UserId, req.TwitId)
	}

	return &pb.Void{}, nil
}
