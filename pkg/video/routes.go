package video

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/micro-tok/gateway/pkg/config"
	"github.com/micro-tok/gateway/pkg/video/routes"
)

func RegisterRouter(r *mux.Router, cfg *config.Config) *ServiceClient {
	client, err := InitServiceClient(cfg)
	if err != nil {
		panic(err)
	}

	svc := &ServiceClient{
		VideoClient: client,
	}

	r.HandleFunc("/upload", svc.UploadVideo).Methods(http.MethodPost)
	r.HandleFunc("/video", svc.GetVideoMetadata).Methods(http.MethodGet)

	return svc
}

func (svc *ServiceClient) UploadVideo(writer http.ResponseWriter, request *http.Request) {
	routes.UploadVideo(writer, request, svc.VideoClient)
}

func (svc *ServiceClient) GetVideoMetadata(writer http.ResponseWriter, request *http.Request) {
	routes.GetVideoMetadata(writer, request, svc.VideoClient)
}
