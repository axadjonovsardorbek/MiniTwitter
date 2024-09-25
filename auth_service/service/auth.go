package service

import (
	"context"

	pb "github.com/axadjonovsardorbek/MiniTwitter/auth_service/genproto/auth"
	st "github.com/axadjonovsardorbek/MiniTwitter/auth_service/storage/postgres"
)

type AuthService struct {
	storage st.Storage
	pb.UnimplementedAuthServiceServer
}

func NewAuthService(storage *st.Storage) *AuthService {
	return &AuthService{storage: *storage}
}

func (s *AuthService) Create(ctx context.Context, req *pb.UserCreateReq) (*pb.TokenRes, error){
	return nil, nil
}
func (s *AuthService) GetById(ctx context.Context, req *pb.ById) (*pb.UserRes, error){
	return nil, nil
}
func (s *AuthService) Update(ctx context.Context, req *pb.UserUpdateReq) (*pb.Void, error){
	return nil, nil
}
func (s *AuthService) Delete(ctx context.Context, req *pb.ById) (*pb.Void, error){
	return nil, nil
}
func (s *AuthService) ChangePassword(ctx context.Context, req *pb.ChangePasswordReq) (*pb.Void, error){
	return nil, nil
}
func (s *AuthService) Login(ctx context.Context, req *pb.LoginReq) (*pb.TokenRes, error){
	return nil, nil
}