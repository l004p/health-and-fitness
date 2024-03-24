package main

import (
	//"fmt"
	"context"
	//"fmt"
	"log"
	"net/http"
	"os"
	"server/api/graph"
	"server/api/graph/model"
	"server/api/middleware"
	"server/api/rest"
	"server/db/pg"
	"server/db/connect"
	//"server/services/user"

	"github.com/99designs/gqlgen/graphql"

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

	ctx := context.Background()

	connection := connect.NewPgConnection(ctx, os.Getenv("DATABASE_URL"))
	defer connection.Close()

	err = connection.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	repo := pg.NewRepository(connection)

	// query := pg.New(connection.Db)

	// users, err := query.GetAllUsers(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, user := range users {
	// 	fmt.Printf("%v\n", user)
	// }

	// fmt.Printf("Printing from repo shit: %d\n", userservice.TestFunction(repo, ctx))

	router := mux.NewRouter()
	protected := router.PathPrefix("/query").Subrouter()
	protected.Use(middleware.AuthMiddleware())

	c := graph.Config{Resolvers: &graph.Resolver{Repo: repo}}
	c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error){
		//fmt.Printf("hasRole Directive %s\n", role.String())
		//userservice.HasRole(repo, ctx, , role.String())
		log.Printf("hasRole Directive %s\n", role)
		
		return next(ctx)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))
	ah :=  rest.NewAuthHandler(repo)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	router.HandleFunc("/login", ah.Login).Methods("POST")
	router.HandleFunc("/register", ah.Register).Methods("POST")

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
