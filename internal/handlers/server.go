package handlers

import (
	"flag"
	"github.com/CvitoyBamp/urlshorter/internal/storage"
	"net/http"
)

type Config struct {
	Address string
}

type Server struct {
	Server  *http.Server
	Storage storage.IStorage
}

func CreateServer() *Server {

	var cfg Config

	flag.StringVar(&cfg.Address, "a", "localhost:8080",
		"An address the server will send metrics")
	flag.Parse()

	return &Server{
		Server: &http.Server{
			Addr: cfg.Address,
		},
		Storage: storage.CreateStorage(),
	}
}

func (server *Server) RunServer() error {
	return http.ListenAndServe(server.Server.Addr, server.ShortURLRouter())
}
