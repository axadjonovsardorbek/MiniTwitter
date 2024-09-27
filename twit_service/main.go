package main


import (
	"log"
	"github.com/axadjonovsardorbek/MiniTwitter/twit_service/config"
	"github.com/axadjonovsardorbek/MiniTwitter/twit_service/service"
	mp "github.com/axadjonovsardorbek/MiniTwitter/twit_service/genproto/twit"
	"github.com/axadjonovsardorbek/MiniTwitter/twit_service/storage/postgres"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cf := config.Load()
	db, err := postgres.NewPostgresStorage(cf)

	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", cf.TWIT_PORT)

	if err != nil {
		log.Fatalf("Failed to listen tcp: %v", err)
	}

	s := grpc.NewServer()

	mp.RegisterFollowServiceServer(s, service.NewFollowService(db))
	mp.RegisterLikeServiceServer(s, service.NewLikeService(db))
	mp.RegisterTwitServiceServer(s, service.NewTwitService(db))

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
