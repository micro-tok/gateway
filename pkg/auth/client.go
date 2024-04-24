package auth

import (
	"errors"
	"net/http"

	authpb "github.com/micro-tok/gateway/pkg/auth/pb"
	"github.com/micro-tok/gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	AuthClient authpb.AuthServiceClient
}

func InitServiceClient(cfg *config.Config) (authpb.AuthServiceClient, error) {
	creds := insecure.NewCredentials()

	c, err := grpc.Dial(cfg.AuthServiceAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	return authpb.NewAuthServiceClient(c), nil
}

func ValidateToken(w http.ResponseWriter, r *http.Request, c authpb.AuthServiceClient) (string, error) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		http.Error(w, "No token provided", http.StatusUnauthorized)
		return "", errors.New("no token provided")
	}

	token := auth[7:]

	validateRequest := &authpb.ValidateUserRequest{
		Token: token,
	}

	res, err := c.ValidateUser(r.Context(), validateRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", err
	}

	return res.UserUuid, nil
}

func GetTags(w http.ResponseWriter, r *http.Request, c authpb.AuthServiceClient) ([]string, error) {
	getMyInterestsRequest := &authpb.GetMyInterestsRequest{
		Token: r.Header.Get("Authorization"),
	}

	res, err := c.GetMyInterests(r.Context(), getMyInterestsRequest)
	if err != nil {
		return nil, err
	}

	return res.Interests, nil
}
