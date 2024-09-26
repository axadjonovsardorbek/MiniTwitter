package service

import (
	"context"
	"log/slog"

	pb "github.com/axadjonovsardorbek/MiniTwitter/twit_service/genproto/twit"
	st "github.com/axadjonovsardorbek/MiniTwitter/twit_service/storage/postgres"
)

type TwitService struct {
	storage st.Storage
	pb.UnimplementedTwitServiceServer
}

func NewTwitService(storage *st.Storage) *TwitService {
	return &TwitService{storage: *storage}
}

func (s *TwitService) Create(ctx context.Context, req *pb.TwitCreateReq) (*pb.Void, error) {
	res, err := s.storage.TwitS.Create(req)
	if err != nil {
		slog.Error("Error creating twit", err)
		return nil, err
	}

	slog.Info("Twit created successfully")
	return res, nil
}
func (s *TwitService) GetById(ctx context.Context, req *pb.ById) (*pb.TwitRes, error) {
	res, err := s.storage.TwitS.GetById(req)
	if err != nil {
		slog.Error("Error getting twit", err)
		return nil, err
	}

	slog.Info("Twit got successfully")
	return res, nil
}
func (s *TwitService) GetAll(ctx context.Context, req *pb.TwitGetAllReq) (*pb.TwitGetAllRes, error) {
	res, err := s.storage.TwitS.GetAll(req)
	if err != nil {
		slog.Error("Error getting twits", err)
		return nil, err
	}

	slog.Info("Twits got successfully")
	return res, nil
}
func (s *TwitService) Update(ctx context.Context, req *pb.TwitUpdateReq) (*pb.Void, error) {
	res, err := s.storage.TwitS.Update(req)
	if err != nil {
		slog.Error("Error updating twit", err)
		return nil, err
	}

	slog.Info("Twit updated successfully")
	return res, nil
}
func (s *TwitService) Delete(ctx context.Context, req *pb.TwitDeleteReq) (*pb.Void, error) {
	res, err := s.storage.TwitS.Delete(req)
	if err != nil {
		slog.Error("Error deleting twit", err)
		return nil, err
	}

	slog.Info("Twit deleted successfully")
	return res, nil
}
