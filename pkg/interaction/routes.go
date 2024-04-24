package interaction

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/micro-tok/gateway/pkg/auth"
	authpb "github.com/micro-tok/gateway/pkg/auth/pb"
	"github.com/micro-tok/gateway/pkg/config"
	"github.com/micro-tok/gateway/pkg/interaction/routes"
)

func RegisterRouter(r *mux.Router, cfg *config.Config, authClient authpb.AuthServiceClient) *ServiceClient {
	svc := &ServiceClient{
		AuthClient:             authClient,
		InteractionServiceAddr: cfg.InteractionServiceAddr,
	}

	r.HandleFunc("/like", svc.Like).Methods(http.MethodPost)
	r.HandleFunc("/like", svc.DeleteLike).Methods(http.MethodDelete)
	r.HandleFunc("/likes/count", svc.CountLikes).Methods(http.MethodGet)
	r.HandleFunc("/likes/find", svc.FindLikes).Methods(http.MethodGet)

	r.HandleFunc("/comment", svc.Comment).Methods(http.MethodPost)
	r.HandleFunc("/comment", svc.DeleteComment).Methods(http.MethodDelete)
	r.HandleFunc("/comments/count", svc.CountComments).Methods(http.MethodGet)
	r.HandleFunc("/comments/find", svc.FindComments).Methods(http.MethodGet)

	return svc
}

func (svc *ServiceClient) Like(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ValidateToken(w, r, svc.AuthClient)
	if err != nil {
		fmt.Println("Error validating token: ", err)
		return
	}

	routes.Like(w, r, userId, svc.InteractionServiceAddr)
}

func (svc *ServiceClient) DeleteLike(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ValidateToken(w, r, svc.AuthClient)
	if err != nil {
		fmt.Println("Error validating token: ", err)
		return
	}

	routes.DeleteLike(w, r, userId, svc.InteractionServiceAddr)
}

func (svc *ServiceClient) CountLikes(w http.ResponseWriter, r *http.Request) {
	routes.CountLikes(w, r, svc.InteractionServiceAddr)
}

func (svc *ServiceClient) FindLikes(w http.ResponseWriter, r *http.Request) {
	routes.FindLikes(w, r, svc.InteractionServiceAddr)
}

func (svc *ServiceClient) Comment(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ValidateToken(w, r, svc.AuthClient)
	if err != nil {
		fmt.Println("Error validating token: ", err)
		return
	}

	routes.Comment(w, r, userId, svc.InteractionServiceAddr)
}

func (svc *ServiceClient) DeleteComment(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ValidateToken(w, r, svc.AuthClient)
	if err != nil {
		fmt.Println("Error validating token: ", err)
		return
	}

	routes.DeleteComment(w, r, userId, svc.InteractionServiceAddr)
}

func (svc *ServiceClient) CountComments(w http.ResponseWriter, r *http.Request) {
	routes.CountComments(w, r, svc.InteractionServiceAddr)
}

func (svc *ServiceClient) FindComments(w http.ResponseWriter, r *http.Request) {
	routes.FindComments(w, r, svc.InteractionServiceAddr)
}
