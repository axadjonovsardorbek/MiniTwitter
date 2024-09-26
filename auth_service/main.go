package main

import (
	"log"
	"log/slog"

	"github.com/axadjonovsardorbek/MiniTwitter/auth_service/api"
	"github.com/axadjonovsardorbek/MiniTwitter/auth_service/api/handler"
	"github.com/axadjonovsardorbek/MiniTwitter/auth_service/config"
	"github.com/axadjonovsardorbek/MiniTwitter/auth_service/minio"
	"github.com/axadjonovsardorbek/MiniTwitter/auth_service/service"
	"github.com/axadjonovsardorbek/MiniTwitter/auth_service/storage/postgres"
)

func main() {
	cf := config.Load()

	conn, err := postgres.NewPostgresStorage(cf)

	if err != nil {
		slog.Error("Failed to connect postgres:", err)
	}

	defer conn.Db.Close()

	us := service.NewAuthService(conn)

	minioClient, err := minio.MinIOConnect(&cf)
	if err != nil {
		slog.Error("Failed to connect to MinIO", err)
		return
	}

	handler := handler.NewHandler(us, minioClient)

	roter := api.NewApi(handler)
	log.Println("Server is running on port ", cf.AUTH_PORT)
	if err := roter.Run(cf.AUTH_PORT); err != nil {
		slog.Error("Error:", err)
	}
}
