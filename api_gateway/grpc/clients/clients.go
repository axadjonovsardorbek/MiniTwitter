package clients

import (
	"github.com/axadjonovsardorbek/MiniTwitter/api_gateway/config"
	bp "github.com/axadjonovsardorbek/MiniTwitter/api_gateway/genproto/twit"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClients struct {
	Follow           bp.FollowServiceClient
	Like    bp.LikeServiceClient
	Twit       bp.TwitServiceClient
}

func NewGrpcClients(cfg *config.Config) (*GrpcClients, error) {
	connB, err := grpc.NewClient(cfg.TWIT_HOST+cfg.TWIT_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &GrpcClients{
		Follow:           bp.NewFollowServiceClient(connB),
		Like:    bp.NewLikeServiceClient(connB),
		Twit:       bp.NewTwitServiceClient(connB),
	}, nil
}
