package main

import (
	"log"
	"net/http"
	"os"
	"server/graphql/generated"
	"server/graphql/resolver"
	"server/repository"
	"server/usecase"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jmoiron/sqlx"
)

func newResolvers(db *sqlx.DB) *resolver.Resolver {
	user := usecase.NewUser(db)

	return resolver.NewResolver(user)
}

func main() {
	db, err := repository.Open()
	if err != nil {
		log.Fatal(err)
	}

	const defaultPort = "1323"
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	resolver := newResolvers(db)
	gc := generated.Config{Resolvers: resolver}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(gc))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
