package service

import (
	"context"
	"log/slog"

	pb "github.com/axadjonovsardorbek/MiniTwitter/twit_service/genproto/twit"
	st "github.com/axadjonovsardorbek/MiniTwitter/twit_service/storage/postgres"
)

type LikeService struct {
	storage st.Storage
	pb.UnimplementedLikeServiceServer
}

func NewLikeService(storage *st.Storage) *LikeService {
	return &LikeService{storage: *storage}
}

func (s *LikeService) Create(ctx context.Context, req *pb.LikeCreateReq) (*pb.Void, error) {
	res, err := s.storage.LikeS.Create(req)
	if err != nil {
		slog.Error("Error creating like", err)
		return nil, err
	}

	slog.Info("Like created successfully")
	return res, nil
}
func (s *LikeService) GetAll(ctx context.Context, req *pb.LikeGetAllReq) (*pb.LikeGetAllRes, error) {
	res, err := s.storage.LikeS.GetAll(req)
	if err != nil {
		slog.Error("Error getting likes", err)
		return nil, err
	}

	slog.Info("Likes got successfully")
	return res, nil
}
func (s *LikeService) Delete(ctx context.Context, req *pb.LikeCreateReq) (*pb.Void, error) {
	res, err := s.storage.LikeS.Delete(req)
	if err != nil {
		slog.Error("Error deleting like", err)
		return nil, err
	}

	slog.Info("Like deleted successfully")
	return res, nil
}
