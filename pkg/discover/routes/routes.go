package routes

import (
	"encoding/json"
	"net/http"

	discoverpb "github.com/micro-tok/gateway/pkg/discover/pb"
)

func DiscoverFeed(w http.ResponseWriter, r *http.Request, client discoverpb.DiscoverServiceClient) {
	res, err := client.DiscoverFeed(r.Context(), &discoverpb.DiscoverFeedRequest{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func DiscoverFeedWithTags(w http.ResponseWriter, r *http.Request, client discoverpb.DiscoverServiceClient, tags []string) {
	res, err := client.DiscoverFeedWithTags(r.Context(), &discoverpb.DiscoverFeedWithTagsRequest{
		Tags: tags,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
