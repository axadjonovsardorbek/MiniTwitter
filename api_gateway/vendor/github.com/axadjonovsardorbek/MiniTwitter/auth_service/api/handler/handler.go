package handler

import (
	"github.com/axadjonovsardorbek/MiniTwitter/auth_service/minio"
	"github.com/axadjonovsardorbek/MiniTwitter/auth_service/service"
)

type Handler struct {
	User  *service.AuthService
	Minio *minio.MinIO
}

func NewHandler(us *service.AuthService, minio *minio.MinIO) *Handler {
	return &Handler{User: us, Minio: minio}
}
