package video

import (
	"github.com/micro-tok/gateway/pkg/config"
	videopb "github.com/micro-tok/gateway/pkg/video/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	VideoClient videopb.VideoServiceClient
}

func InitServiceClient(cfg *config.Config) (videopb.VideoServiceClient, error) {
	creds := insecure.NewCredentials()

	c, err := grpc.Dial(cfg.VideoServiceAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	return videopb.NewVideoServiceClient(c), nil
}
