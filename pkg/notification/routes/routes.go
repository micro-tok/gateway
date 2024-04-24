package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	notificationpb "github.com/micro-tok/gateway/pkg/notification/pb"
)

type TokenReq struct {
	token string
}

func RegisterToken(w http.ResponseWriter, r *http.Request, client notificationpb.NotificationServiceClient, userId string) {
	var tokenReq TokenReq
	err := json.NewDecoder(r.Body).Decode(&tokenReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = client.RegisterToken(r.Context(), &notificationpb.RegisterTokenRequest{
		Uuid:  userId,
		Token: tokenReq.token,
	})

	if err != nil {
		fmt.Println("Error registering token: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func RemoveToken(w http.ResponseWriter, r *http.Request, client notificationpb.NotificationServiceClient, userId string) {
	var tokenReq TokenReq
	err := json.NewDecoder(r.Body).Decode(&tokenReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = client.RemoveToken(r.Context(), &notificationpb.RemoveTokenRequest{
		Uuid:  userId,
		Token: tokenReq.token,
	})

	if err != nil {
		fmt.Println("Error removing token: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func RemoveUserTokens(w http.ResponseWriter, r *http.Request, client notificationpb.NotificationServiceClient, userId string) {
	_, err := client.RemoveUserTokens(r.Context(), &notificationpb.RemoveUserTokensRequest{
		Uuid: userId,
	})

	if err != nil {
		fmt.Println("Error removing user tokens: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetUserTokens(w http.ResponseWriter, r *http.Request, client notificationpb.NotificationServiceClient, userId string) {
	res, err := client.GetUserTokens(r.Context(), &notificationpb.GetUserTokensRequest{
		Uuid: userId,
	})

	if err != nil {
		fmt.Println("Error getting user tokens: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetAllTokens(w http.ResponseWriter, r *http.Request, client notificationpb.NotificationServiceClient) {
	res, err := client.GetAllTokens(r.Context(), &notificationpb.GetAllTokensRequest{})

	if err != nil {
		fmt.Println("Error getting all tokens: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
