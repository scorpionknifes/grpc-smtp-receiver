package main

import "time"

// Config file struct
type Config struct {
	SMTPPort string
	Timeout  time.Duration
	Server   []Server
}

// Server config struct
type Server struct {
	EmailDomain string
	GRPCServer  string
}
