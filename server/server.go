package main

import (
	//"fmt"
	"log"
	"net/http"
	"os"
	"server/api/graph"
	"server/api/middleware"
	"context"
	"github.com/99designs/gqlgen/graphql"
	"server/api/graph/model"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load("../.env")
		if err != nil {
			log.Fatal(err)
		}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := mux.NewRouter()
	router.Use(middleware.AuthMiddleware())

	c := graph.Config{Resolvers: &graph.Resolver{}}
	c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error){
		//fmt.Printf("hasRole Directive %s\n", role.String())
		log.Printf("hasRole Directive %s\n", role)
		return next(ctx)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
