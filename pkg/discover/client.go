package discover

import (
	authpb "github.com/micro-tok/gateway/pkg/auth/pb"
	"github.com/micro-tok/gateway/pkg/config"
	discoverpb "github.com/micro-tok/gateway/pkg/discover/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	DiscoverClient discoverpb.DiscoverServiceClient
	AuthClient     authpb.AuthServiceClient
}

func InitServiceClient(cfg *config.Config) (discoverpb.DiscoverServiceClient, error) {
	creds := insecure.NewCredentials()

	c, err := grpc.Dial(cfg.NotificationServiceAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	return discoverpb.NewDiscoverServiceClient(c), nil
}
