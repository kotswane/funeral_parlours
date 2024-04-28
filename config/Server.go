package config

import "os"

type Server struct {
	Host string
	Port string
}

func NewServer() Server {
	return Server{
		Host: os.Getenv("Host"),
		Port: os.Getenv("Port"),
	}
}
