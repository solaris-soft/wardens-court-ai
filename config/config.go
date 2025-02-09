package config

import (
	"net/http"
	"os"
	"time"
)

// Config stores server configuration settings
type Config struct {
	Server *http.Server
	DBPath string
}

// NewAppConfig returns a new app configuration with default settings
func NewAppConfig() *Config {
	server := http.Server{
		Addr:         os.Getenv("ADDR"),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return &Config{
		Server: &server,
		DBPath: os.Getenv("DB_PATH"),
	}
}

// GetServeAddr returns the serve address and port number from the config
func (c *Config) GetServeAddr() string {
	return c.Server.Addr
}
