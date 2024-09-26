package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	pb "github.com/axadjonovsardorbek/MiniTwitter/twit_service/genproto/twit"
	"github.com/google/uuid"
)

type FollowRepo struct {
	db *sql.DB
}

func NewFollowRepo(db *sql.DB) *FollowRepo {
	return &FollowRepo{db: db}
}

func (r *FollowRepo) Create(req *pb.FollowCreateReq) (*pb.Void, error) {
	id := uuid.New().String()

	query := `
	INSERT INTO follows(
		id,
		follower_id,
		followed_id
	) VALUES($1, $2, $3)
	`

	_, err := r.db.Exec(query, id, req.FollowerId, req.FollowedId)

	if err != nil {
		fmt.Println("error while following")
		return nil, errors.New("error while following")
	}

	return nil, nil
}
func (r *FollowRepo) GetAll(req *pb.FollowGetAllReq) (*pb.FollowGetAllRes, error) {
	res := &pb.FollowGetAllRes{}

	query := `
	SELECT
		COUNT(id) OVER () AS total_count,
		id,
		follower_id,
		followed_id
	FROM
		follows
	WHERE 
		deleted_at = 0
	`

	var args []interface{}
	var conditions []string

	if req.FollowerId != "" && req.FollowerId != "string" {
		conditions = append(conditions, fmt.Sprintf("follower_id = $%d", len(args)+1))
		args = append(args, req.FollowerId)
	}

	if req.FollowedId != "" && req.FollowedId != "string" {
		conditions = append(conditions, fmt.Sprintf("followed_id = $%d", len(args)+1))
		args = append(args, req.FollowedId)
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("follows list is empty")
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		follow := pb.FollowRes{}
		var count int32

		err := rows.Scan(
			&count,
			&follow.Id,
			&follow.FollowerId,
			&follow.FollowedId,
		)
		if err != nil {
			return nil, err
		}

		res.Follows = append(res.Follows, &follow)
		res.Count = count
	}

	return res, nil
}
func (r *FollowRepo) Delete(req *pb.FollowCreateReq) (*pb.Void, error) {

	query := `
	UPDATE 
		follows
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		follower_id = $1
	AND 
		followed_id = $2
	AND
		deleted_at = 0
	`

	res, err := r.db.Exec(query, req.FollowerId, req.FollowedId)

	if err != nil {
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("follow with follower_id %s and followed_id %s not found", req.FollowerId, req.FollowedId)
	}

	return &pb.Void{}, nil
}
