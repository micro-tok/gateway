package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	authpb "github.com/micro-tok/gateway/pkg/auth/pb"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func Login(w http.ResponseWriter, r *http.Request, c authpb.AuthServiceClient) {
	var loginReq LoginReq
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	loginRequest := &authpb.LoginRequest{
		Email:    loginReq.Email,
		Password: loginReq.Password,
	}

	res, err := c.Login(r.Context(), loginRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func Register(w http.ResponseWriter, r *http.Request, c authpb.AuthServiceClient) {
	var registerReq RegisterReq
	err := json.NewDecoder(r.Body).Decode(&registerReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	registerRequest := &authpb.RegisterRequest{
		Email:    registerReq.Email,
		Password: registerReq.Password,
		Username: registerReq.Username,
	}

	res, err := c.Register(r.Context(), registerRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func RefreshToken(w http.ResponseWriter, r *http.Request, c authpb.AuthServiceClient) {
	refreshTokenRequest := &authpb.RefreshTokenRequest{
		RefreshToken: r.Header.Get("Authorization"),
	}

	res, err := c.RefreshToken(r.Context(), refreshTokenRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func ValidateUser(next http.HandlerFunc, c authpb.AuthServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		validateUserRequest := &authpb.ValidateUserRequest{
			Token: r.Header.Get("Authorization"),
		}

		res, err := c.ValidateUser(r.Context(), validateUserRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)

		next(w, r)
	}
}

func GetMe(w http.ResponseWriter, r *http.Request, c authpb.AuthServiceClient) {
	getMeRequest := &authpb.GetMeRequest{
		Token: r.Header.Get("Authorization"),
	}

	res, err := c.GetMe(r.Context(), getMeRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func UpdateMyInterests(w http.ResponseWriter, r *http.Request, c authpb.AuthServiceClient) {
	var interests []string
	err := json.NewDecoder(r.Body).Decode(&interests)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateMyInterestsRequest := &authpb.UpdateMyInterestsRequest{
		Token:     r.Header.Get("Authorization"),
		Interests: interests,
	}

	res, err := c.UpdateMyInterests(r.Context(), updateMyInterestsRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetMyInterests(w http.ResponseWriter, r *http.Request, c authpb.AuthServiceClient) {
	getMyInterestsRequest := &authpb.GetMyInterestsRequest{
		Token: r.Header.Get("Authorization"),
	}

	res, err := c.GetMyInterests(r.Context(), getMyInterestsRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetPublicUser(w http.ResponseWriter, r *http.Request, c authpb.AuthServiceClient) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	getPublicUserRequest := &authpb.GetPublicUserRequest{
		UserUuid: id,
	}

	res, err := c.GetPublicUser(r.Context(), getPublicUserRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request, c authpb.AuthServiceClient) {
	getAllUsersRequest := &authpb.GetAllUsersRequest{}

	res, err := c.GetAllUsers(r.Context(), getAllUsersRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
