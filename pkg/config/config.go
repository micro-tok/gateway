package config

import "os"

type Config struct {
	Port                    string
	VideoServiceAddr        string
	AuthServiceAddr         string
	NotificationServiceAddr string
	DiscoverServiceAddr     string
	InteractionServiceAddr  string
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	return &Config{
		Port:                    port,
		VideoServiceAddr:        os.Getenv("VIDEO_SERVICE_ADDR"),
		AuthServiceAddr:         os.Getenv("AUTH_SERVICE_ADDR"),
		NotificationServiceAddr: os.Getenv("NOTIFICATION_SERVICE_ADDR"),
		DiscoverServiceAddr:     os.Getenv("DISCOVER_SERVICE_ADDR"),
		InteractionServiceAddr:  os.Getenv("INTERACTION_SERVICE_ADDR"),
	}
}
