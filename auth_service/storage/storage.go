package storage

import (
	pb "github.com/axadjonovsardorbek/MiniTwitter/auth_service/genproto"
)

type AuthI interface{
    Create(*pb.UserCreateReq) (*pb.TokenRes, error)
    GetById(*pb.ById) (*pb.UserRes, error)
    Update(*pb.UserUpdateReq) (*pb.Void, error)
    Delete(*pb.ById) (*pb.Void, error)
    ChangePassword(*pb.ChangePasswordReq) (*pb.Void, error)
    Login(*pb.LoginReq) (*pb.TokenRes, error)
}