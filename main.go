package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"github.com/micro-tok/gateway/pkg/auth"
	"github.com/micro-tok/gateway/pkg/config"
	"github.com/micro-tok/gateway/pkg/video"
	"github.com/rs/cors"
)

func main() {
	cfg := config.LoadConfig()

	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://*.vercel.app", "http://localhost:3000/", "https://*.vercel.app/", "http://*.kaloyan.tech/", "https://*.kaloyan.tech", "*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})

	authSvc := auth.RegisterRouter(r, cfg)

	video.RegisterRouter(r, cfg)

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler := c.Handler(r)

	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler))
}
