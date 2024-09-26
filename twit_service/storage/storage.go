package storage

import (
	pb "github.com/axadjonovsardorbek/MiniTwitter/twit_service/genproto/twit"
)

type FollowI interface {
	Create(*pb.FollowCreateReq) (*pb.Void, error)
	GetAll(*pb.FollowGetAllReq) (*pb.FollowGetAllRes, error)
	Delete(*pb.FollowCreateReq) (*pb.Void, error)
}

type LikeI interface {
	Create(*pb.LikeCreateReq) (*pb.Void, error)
	GetAll(*pb.LikeGetAllReq) (*pb.LikeGetAllRes, error)
	Delete(*pb.LikeCreateReq) (*pb.Void, error)
}

type TwitI interface {
	Create(*pb.TwitCreateReq) (*pb.Void, error)
	GetById(*pb.ById) (*pb.TwitRes, error)
	GetAll(*pb.TwitGetAllReq) (*pb.TwitGetAllRes, error)
	Update(*pb.TwitUpdateReq) (*pb.Void, error)
	Delete(*pb.TwitDeleteReq) (*pb.Void, error)
}
