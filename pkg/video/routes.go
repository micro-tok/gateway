package video

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/micro-tok/gateway/pkg/auth"
	authpb "github.com/micro-tok/gateway/pkg/auth/pb"
	"github.com/micro-tok/gateway/pkg/config"
	"github.com/micro-tok/gateway/pkg/video/routes"
)

func RegisterRouter(r *mux.Router, cfg *config.Config, authClient authpb.AuthServiceClient) *ServiceClient {
	client, err := InitServiceClient(cfg)
	if err != nil {
		panic(err)
	}

	svc := &ServiceClient{
		VideoClient: client,
		AuthClient:  authClient,
	}

	r.HandleFunc("/upload", svc.UploadVideo).Methods(http.MethodPost)
	r.HandleFunc("/video", svc.GetVideoMetadata).Methods(http.MethodGet)

	return svc
}

func (svc *ServiceClient) UploadVideo(writer http.ResponseWriter, request *http.Request) {
	userId, err := auth.ValidateToken(writer, request, svc.AuthClient)
	if err != nil {
		return
	}
	routes.UploadVideo(writer, request, svc.VideoClient, userId)
}

func (svc *ServiceClient) GetVideoMetadata(writer http.ResponseWriter, request *http.Request) {
	routes.GetVideoMetadata(writer, request, svc.VideoClient)
}
