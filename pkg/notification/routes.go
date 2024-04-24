package notification

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/micro-tok/gateway/pkg/auth"
	authpb "github.com/micro-tok/gateway/pkg/auth/pb"
	"github.com/micro-tok/gateway/pkg/config"
	"github.com/micro-tok/gateway/pkg/notification/routes"
)

func RegisterRouter(r *mux.Router, cfg *config.Config, authClient authpb.AuthServiceClient) *ServiceClient {
	client, err := InitServiceClient(cfg)
	if err != nil {
		panic(err)
	}

	svc := &ServiceClient{
		NotificationClient: client,
		AuthClient:         authClient,
	}

	r.HandleFunc("/token", svc.RegisterToken).Methods(http.MethodPost)
	r.HandleFunc("/token", svc.RemoveToken).Methods(http.MethodDelete)
	r.HandleFunc("/tokens", svc.RemoveUserTokens).Methods(http.MethodDelete)
	r.HandleFunc("/tokens", svc.GetTokens).Methods(http.MethodGet)
	r.HandleFunc("/tokens-all/all", svc.GetAllTokens).Methods(http.MethodGet)
	return svc
}

func (svc *ServiceClient) RegisterToken(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ValidateToken(w, r, svc.AuthClient)
	if err != nil {
		fmt.Println("Error validating token: ", err)
		return
	}

	routes.RegisterToken(w, r, svc.NotificationClient, userId)
}

func (svc *ServiceClient) RemoveToken(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ValidateToken(w, r, svc.AuthClient)
	if err != nil {
		fmt.Println("Error validating token: ", err)
		return
	}

	routes.RemoveToken(w, r, svc.NotificationClient, userId)
}

func (svc *ServiceClient) RemoveUserTokens(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ValidateToken(w, r, svc.AuthClient)
	if err != nil {
		fmt.Println("Error validating token: ", err)
		return
	}

	routes.RemoveUserTokens(w, r, svc.NotificationClient, userId)
}

func (svc *ServiceClient) GetTokens(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ValidateToken(w, r, svc.AuthClient)
	if err != nil {
		fmt.Println("Error validating token: ", err)
		return
	}

	routes.GetUserTokens(w, r, svc.NotificationClient, userId)
}

func (svc *ServiceClient) GetAllTokens(w http.ResponseWriter, r *http.Request) {
	routes.GetAllTokens(w, r, svc.NotificationClient)
}
