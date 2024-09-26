package service

import (
	"context"
	"log/slog"

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

func (s *AuthService) Create(ctx context.Context, req *pb.UserCreateReq) (*pb.TokenRes, error) {
	res, err := s.storage.AuthS.Create(req)
	if err != nil {
		slog.Error("Error creating user", err)
		return nil, err
	}

	slog.Info("User created successfully")
	return res, nil
}
func (s *AuthService) GetById(ctx context.Context, req *pb.ById) (*pb.UserRes, error) {
	res, err := s.storage.AuthS.GetById(req)
	if err != nil {
		slog.Error("Error getting user", err)
		return nil, err
	}

	slog.Info("User got successfully")
	return res, nil
}
func (s *AuthService) Update(ctx context.Context, req *pb.UserUpdateReq) (*pb.Void, error) {
	res, err := s.storage.AuthS.Update(req)
	if err != nil {
		slog.Error("Error updating user", err)
		return nil, err
	}

	slog.Info("User updated successfully")
	return res, nil
}
func (s *AuthService) Delete(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	res, err := s.storage.AuthS.Delete(req)
	if err != nil {
		slog.Error("Error deleting user", err)
		return nil, err
	}

	slog.Info("User deleted successfully")
	return res, nil
}
func (s *AuthService) ChangePassword(ctx context.Context, req *pb.ChangePasswordReq) (*pb.Void, error) {
	res, err := s.storage.AuthS.ChangePassword(req)
	if err != nil {
		slog.Error("Error changing password", err)
		return nil, err
	}

	slog.Info("Password changed successfully")
	return res, nil
}
func (s *AuthService) Login(ctx context.Context, req *pb.LoginReq) (*pb.TokenRes, error) {
	res, err := s.storage.AuthS.Login(req)
	if err != nil {
		slog.Error("Error login user", err)
		return nil, err
	}

	slog.Info("User logged successfully")
	return res, nil
}
