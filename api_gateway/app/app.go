package app

import (
	"log"

	"github.com/axadjonovsardorbek/MiniTwitter/api_gateway/api"
	_ "github.com/axadjonovsardorbek/MiniTwitter/api_gateway/api/docs"
	"github.com/axadjonovsardorbek/MiniTwitter/api_gateway/api/handler"
	"github.com/axadjonovsardorbek/MiniTwitter/api_gateway/config"
	"github.com/axadjonovsardorbek/MiniTwitter/api_gateway/grpc/clients"
	"golang.org/x/exp/slog"
)

func Run(cfg *config.Config) {

	services, err := clients.NewGrpcClients(cfg)
	if err != nil {
		log.Fatalf("error while connecting clients. err: %s", err.Error())
	}

	if err != nil {
		log.Fatal(err)
		return
	}

	h := handler.NewHandler(services)

	router := api.NewApi(h)
	if err := router.Run(cfg.API_PORT); err != nil {
		slog.Error("Failed to start API Gateway: %v", err)
	}
}
