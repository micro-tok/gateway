package notification

import (
	authpb "github.com/micro-tok/gateway/pkg/auth/pb"
	"github.com/micro-tok/gateway/pkg/config"
	notificationpb "github.com/micro-tok/gateway/pkg/notification/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	NotificationClient notificationpb.NotificationServiceClient
	AuthClient         authpb.AuthServiceClient
}

func InitServiceClient(cfg *config.Config) (notificationpb.NotificationServiceClient, error) {
	creds := insecure.NewCredentials()

	c, err := grpc.Dial(cfg.NotificationServiceAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	return notificationpb.NewNotificationServiceClient(c), nil
}
