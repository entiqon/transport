package config

import "time"

// API defines configuration for HTTP-based transports.
type API struct {
	BaseURL string
	Timeout time.Duration
}

// SFTP defines configuration for SFTP-based transports.
type SFTP struct {
	Host string
	Port int
	User string
}
