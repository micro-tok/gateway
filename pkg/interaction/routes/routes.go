package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type CommentDtoReq struct {
	UUID    string `json:"uuid"`
	UPID    string `json:"upid"`
	PUUID   string `json:"puuid"`
	Content string `json:"content"`
}

type LikeDtoReq struct {
	UUID  string `json:"uuid"`
	UPID  string `json:"upid"`
	PUUID string `json:"puuid"`
}

type UPIDDtoReq struct {
	UPID string `json:"upid"`
}

func Like(w http.ResponseWriter, r *http.Request, userId string, interactionServiceAddr string) {
	var likeDtoReq LikeDtoReq
	err := json.NewDecoder(r.Body).Decode(&likeDtoReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	likeDtoReq.UUID = userId

	// http req to interaction service to like
	reqString := fmt.Sprintf(`{"uuid": "%s", "upid": "%s", "puuid": "%s"}`, likeDtoReq.UUID, likeDtoReq.UPID, likeDtoReq.PUUID)
	req, err := http.NewRequest(http.MethodPost, interactionServiceAddr+"/api/like/save", strings.NewReader(reqString))
	if err != nil {
		fmt.Println("Error liking: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error liking: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func Comment(w http.ResponseWriter, r *http.Request, userId string, interactionServiceAddr string) {
	var commentDtoReq CommentDtoReq
	err := json.NewDecoder(r.Body).Decode(&commentDtoReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	commentDtoReq.UUID = userId

	// http req to interaction service to comment
	reqString := fmt.Sprintf(`{"uuid": "%s", "upid": "%s", "puuid": "%s", "content": "%s"}`, commentDtoReq.UUID, commentDtoReq.UPID, commentDtoReq.PUUID, commentDtoReq.Content)
	req, err := http.NewRequest(http.MethodPost, interactionServiceAddr+"/api/comment/save", strings.NewReader(reqString))
	if err != nil {
		fmt.Println("Error commenting: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error commenting: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func DeleteLike(w http.ResponseWriter, r *http.Request, userId string, interactionServiceAddr string) {
	var likeDtoReq LikeDtoReq
	err := json.NewDecoder(r.Body).Decode(&likeDtoReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	likeDtoReq.UUID = userId

	// http req to interaction service to delete like
	reqString := fmt.Sprintf(`{"uuid": "%s", "upid": "%s", "puuid": "%s"}`, likeDtoReq.UUID, likeDtoReq.UPID, likeDtoReq.PUUID)
	req, err := http.NewRequest(http.MethodDelete, interactionServiceAddr+"/api/like/delete", strings.NewReader(reqString))
	if err != nil {
		fmt.Println("Error deleting like: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error deleting like: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func DeleteComment(w http.ResponseWriter, r *http.Request, userId string, interactionServiceAddr string) {
	var commentDtoReq CommentDtoReq
	err := json.NewDecoder(r.Body).Decode(&commentDtoReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	commentDtoReq.UUID = userId

	// http req to interaction service to delete comment
	reqString := fmt.Sprintf(`{"uuid": "%s", "upid": "%s", "puuid": "%s"}`, commentDtoReq.UUID, commentDtoReq.UPID, commentDtoReq.PUUID)
	req, err := http.NewRequest(http.MethodDelete, interactionServiceAddr+"/api/comment/delete", strings.NewReader(reqString))
	if err != nil {
		fmt.Println("Error deleting comment: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error deleting comment: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func CountLikes(w http.ResponseWriter, r *http.Request, interactionServiceAddr string) {
	upid := r.URL.Query().Get("upid")
	if upid == "" {
		http.Error(w, "upid is required", http.StatusBadRequest)
		return
	}

	var upidDtoReq UPIDDtoReq
	upidDtoReq.UPID = upid

	// http req to interaction service to count likes
	req, err := http.NewRequest(http.MethodGet, interactionServiceAddr+"/api/like/countAllByPostId", strings.NewReader(fmt.Sprintf(`{"upid": "%s"}`, upidDtoReq.UPID)))
	if err != nil {
		fmt.Println("Error counting likes: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error counting likes: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func CountComments(w http.ResponseWriter, r *http.Request, interactionServiceAddr string) {
	upid := r.URL.Query().Get("upid")
	if upid == "" {
		http.Error(w, "upid is required", http.StatusBadRequest)
		return
	}

	var upidDtoReq UPIDDtoReq
	upidDtoReq.UPID = upid

	// http req to interaction service to count comments
	req, err := http.NewRequest(http.MethodGet, interactionServiceAddr+"/api/comment/countAllByPostId", strings.NewReader(fmt.Sprintf(`{"upid": "%s"}`, upidDtoReq.UPID)))
	if err != nil {
		fmt.Println("Error counting comments: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error counting comments: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func FindLikes(w http.ResponseWriter, r *http.Request, interactionServiceAddr string) {
	upid := r.URL.Query().Get("upid")
	if upid == "" {
		http.Error(w, "upid is required", http.StatusBadRequest)
		return
	}

	var upidDtoReq UPIDDtoReq
	upidDtoReq.UPID = upid

	// http req to interaction service to find likes
	req, err := http.NewRequest(http.MethodGet, interactionServiceAddr+"/api/like/findAllByPostId", strings.NewReader(fmt.Sprintf(`{"upid": "%s"}`, upidDtoReq.UPID)))
	if err != nil {
		fmt.Println("Error finding likes: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error finding likes: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func FindComments(w http.ResponseWriter, r *http.Request, interactionServiceAddr string) {
	upid := r.URL.Query().Get("upid")
	if upid == "" {
		http.Error(w, "upid is required", http.StatusBadRequest)
		return
	}

	var upidDtoReq UPIDDtoReq
	upidDtoReq.UPID = upid

	// http req to interaction service to find comments
	req, err := http.NewRequest(http.MethodGet, interactionServiceAddr+"/api/comment/findAllByPostId", strings.NewReader(fmt.Sprintf(`{"upid": "%s"}`, upidDtoReq.UPID)))
	if err != nil {
		fmt.Println("Error finding comments: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error finding comments: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
