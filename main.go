package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"http-go-sandbox/graph"
	"http-go-sandbox/middlewares"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := chi.NewRouter()

	// middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middlewares.Auth())
	r.Use(middleware.Recoverer)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	r.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":"+port, r))
}
