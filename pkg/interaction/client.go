package interaction

import (
	authpb "github.com/micro-tok/gateway/pkg/auth/pb"
)

type ServiceClient struct {
	AuthClient             authpb.AuthServiceClient
	InteractionServiceAddr string
}
