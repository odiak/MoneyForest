package main

import (
	"log"
	"net/http"
	"time"

	"./apirouter"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Mount("/api", apirouter.NewRouter())

	log.Println("starting server")
	http.ListenAndServe("127.0.0.1:3333", r)
}
