package main

import (
	"github.com/CvitoyBamp/urlshorter/internal/handlers"
	"log"
)

func main() {
	s := handlers.CreateServer()
	err := s.RunServer()
	if err != nil {
		log.Fatal("can't start server")
	}
}
