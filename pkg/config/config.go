package config

import "os"

type Config struct {
	Port             string
	VideoServiceAddr string
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		Port:             port,
		VideoServiceAddr: os.Getenv("VIDEO_SERVICE_ADDR"),
	}
}
