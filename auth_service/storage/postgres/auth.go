package postgres

import (
	"database/sql"

	pb "github.com/axadjonovsardorbek/MiniTwitter/tree/main/auth_service/genproto/auth"

	"github.com/go-redis/redis/v8"
)

type AuthRepo struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewAuthRepo(db *sql.DB, rdb *redis.Client) *AuthRepo {
	return &AuthRepo{db: db, rdb: rdb}
}

func (r *AuthRepo) Create(req *pb.UserCreateReq) (*pb.TokenRes, error) {
	return nil, nil
}
func (r *AuthRepo) GetById(req *pb.ById) (*pb.UserRes, error) {
	return nil, nil
}
func (r *AuthRepo) Update(req *pb.UserUpdateReq) (*pb.Void, error) {
	return nil, nil
}
func (r *AuthRepo) Delete(req *pb.ById) (*pb.Void, error) {
	return nil, nil
}
func (r *AuthRepo) ChangePassword(req *pb.ChangePasswordReq) (*pb.Void, error) {
	return nil, nil
}
func (r *AuthRepo) Login(req *pb.LoginReq) (*pb.TokenRes, error) {
	return nil, nil
}
