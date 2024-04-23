package auth

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/micro-tok/gateway/pkg/auth/routes"
	"github.com/micro-tok/gateway/pkg/config"
)

func RegisterRouter(r *mux.Router, cfg *config.Config) *ServiceClient {
	client, err := InitServiceClient(cfg)
	if err != nil {
		panic(err)
	}

	svc := &ServiceClient{
		AuthClient: client,
	}

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		routes.Login(w, r, svc.AuthClient)
	}).Methods(http.MethodPost)

	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		routes.Register(w, r, svc.AuthClient)
	}).Methods(http.MethodPost)

	r.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		routes.GetMe(w, r, svc.AuthClient)
	}).Methods(http.MethodGet)

	r.HandleFunc("/interests", func(w http.ResponseWriter, r *http.Request) {
		routes.UpdateMyInterests(w, r, svc.AuthClient)
	}).Methods(http.MethodPut)

	r.HandleFunc("/interests", func(w http.ResponseWriter, r *http.Request) {
		routes.GetMyInterests(w, r, svc.AuthClient)
	}).Methods(http.MethodGet)

	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		routes.GetPublicUser(w, r, svc.AuthClient)
	}).Methods(http.MethodGet)

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		routes.GetAllUsers(w, r, svc.AuthClient)
	}).Methods(http.MethodGet)

	r.HandleFunc("/refresh", func(w http.ResponseWriter, r *http.Request) {
		routes.RefreshToken(w, r, svc.AuthClient)
	}).Methods(http.MethodPost)

	return svc
}
