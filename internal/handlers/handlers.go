package handlers

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"io"
	"log"
	"net/http"
	"net/url"
)

func (server *Server) ShortURLRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	//c := cors.New(cors.Options{
	//	AllowedOrigins:     []string{"https://*", "http://*"},
	//	AllowedMethods:     []string{"GET", "POST", "OPTION"},
	//	AllowedHeaders:     []string{"Content-Type", "Location"},
	//	OptionsPassthrough: true,
	//})
	//r.Use(c.Handler)

	r.Route("/", func(r chi.Router) {
		r.Post("/", server.shorterCreatorHandler)
		r.Get("/{id}", server.shorterHandler)
	})

	return r
}

func (server *Server) shorterCreatorHandler(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	s, bodyErr := io.ReadAll(req.Body)
	if bodyErr != nil {
		log.Print("can't read body from request")
	}

	uri, uriErr := url.Parse(string(s))
	if uriErr != nil {
		log.Print("can't parse url")
		http.Error(res, "can't parse url", http.StatusBadRequest)
		return
	}

	val, addErr := server.Storage.AddURL(uri.String())
	if addErr != nil {
		http.Error(res, fmt.Sprintf("Error while creating short url: %s", addErr), http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusCreated)
	res.Header().Set("Content-Type", "text/plain")
	res.Write([]byte(val))
	log.Printf("successful response for shorter %s for address %s", val, uri.String())
}

func (server *Server) shorterHandler(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	val, err := server.Storage.GetURL(id)
	if err != nil {
		http.Error(res, fmt.Sprintf("can't resolve url: %s", err), http.StatusBadRequest)
	}

	if val != "" {
		res.Header().Set("Location", val)
		res.WriteHeader(http.StatusTemporaryRedirect)
		log.Printf("Successfully set header Location as %s in response", val)
		return
	}

	http.Error(res, fmt.Sprintf("can't resolve url: %s", err), http.StatusBadRequest)
}
