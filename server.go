package main

import (
	"log"
	"net/http"
	"os"

	"graphql.example/internal/auth"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"graphql.example/graph"
	hackernews "graphql.example/graph/generated"
	database "graphql.example/internal/pkg/db/migrations/mysql"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	database.InitDB()
	database.Migrate()
	server := handler.GraphQL(hackernews.NewExecutableSchema(hackernews.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
