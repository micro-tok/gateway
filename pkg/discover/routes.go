package discover

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/micro-tok/gateway/pkg/auth"
	authpb "github.com/micro-tok/gateway/pkg/auth/pb"
	"github.com/micro-tok/gateway/pkg/config"
	"github.com/micro-tok/gateway/pkg/discover/routes"
)

func RegisterRouter(r *mux.Router, cfg *config.Config, authClient authpb.AuthServiceClient) *ServiceClient {
	client, err := InitServiceClient(cfg)
	if err != nil {
		panic(err)
	}

	svc := &ServiceClient{
		DiscoverClient: client,
		AuthClient:     authClient,
	}

	r.HandleFunc("/discover", svc.Discover).Methods(http.MethodGet)

	return svc
}

func (svc *ServiceClient) Discover(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") == "" {
		routes.DiscoverFeed(w, r, svc.DiscoverClient)
		return
	}

	tags, err := auth.GetTags(w, r, svc.AuthClient)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	routes.DiscoverFeedWithTags(w, r, svc.DiscoverClient, tags)
}
