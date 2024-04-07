package routes

import (
	"encoding/json"
	"io"
	"net/http"

	videopb "github.com/micro-tok/gateway/pkg/video/pb"
)

func UploadVideo(w http.ResponseWriter, r *http.Request, c videopb.VideoServiceClient) {
	uploadVideoRequest := &videopb.UploadVideoRequest{}

	file, _, err := r.FormFile("video")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uploadVideoRequest.Video = fileBytes
	uploadVideoRequest.Title = r.FormValue("title")
	uploadVideoRequest.Description = r.FormValue("description")
	uploadVideoRequest.Tags = append(uploadVideoRequest.Tags, r.Form["tags"]...)

	res, err := c.UploadVideo(r.Context(), uploadVideoRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

func GetVideoMetadata(w http.ResponseWriter, r *http.Request, c videopb.VideoServiceClient) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	getVideoMetadataRequest := &videopb.GetVideoMetadataRequest{
		Id: id,
	}

	res, err := c.GetVideoMetadata(r.Context(), getVideoMetadataRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
