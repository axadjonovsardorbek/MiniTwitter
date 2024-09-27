package handler

import (
	"github.com/axadjonovsardorbek/MiniTwitter/api_gateway/grpc/clients"
)

type Handler struct {
	srvs *clients.GrpcClients
}

func NewHandler(srvs *clients.GrpcClients) *Handler {
	return &Handler{srvs: srvs}
}
