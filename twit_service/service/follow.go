package service

import (
	"context"
	"log/slog"

	pb "github.com/axadjonovsardorbek/MiniTwitter/twit_service/genproto/twit"
	st "github.com/axadjonovsardorbek/MiniTwitter/twit_service/storage/postgres"
)

type FollowService struct {
	storage st.Storage
	pb.UnimplementedFollowServiceServer
}

func NewFollowService(storage *st.Storage) *FollowService {
	return &FollowService{storage: *storage}
}

func (s *FollowService) Create(ctx context.Context, req *pb.FollowCreateReq) (*pb.Void, error) {
	res, err := s.storage.FollowS.Create(req)
	if err != nil {
		slog.Error("Error creating follow", err)
		return nil, err
	}

	slog.Info("Follow created successfully")
	return res, nil
}
func (s *FollowService) GetAll(ctx context.Context, req *pb.FollowGetAllReq) (*pb.FollowGetAllRes, error) {
	res, err := s.storage.FollowS.GetAll(req)
	if err != nil {
		slog.Error("Error getting follows", err)
		return nil, err
	}

	slog.Info("Follows got successfully")
	return res, nil
}
func (s *FollowService) Delete(ctx context.Context, req *pb.FollowCreateReq) (*pb.Void, error) {
	res, err := s.storage.FollowS.Delete(req)
	if err != nil {
		slog.Error("Error deleting follow", err)
		return nil, err
	}

	slog.Info("Follow deleted successfully")
	return res, nil
}
