package handler

import (
	"github.com/axadjonovsardorbek/MiniTwitter/auth_service/service"
)

type Handler struct {
	User *service.AuthService
}

func NewHandler(us *service.AuthService) *Handler {
	return &Handler{User: us}
}
