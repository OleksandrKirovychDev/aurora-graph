package main

import (
	"aurora-graph/gateway/config"
	"aurora-graph/gateway/graph"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	server, err := graph.NewGraphQLServer(config.AccountUrl)

    if err != nil {
        log.Fatal(err)
    }
	defer server.AccountClient.Close()


  	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: server},
		),
	)

	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

    http.Handle("/", playground.Handler("GraphQL playground", "/query"))
    http.Handle("/query", srv)

    log.Println("GraphQL gateway running at http://localhost:8000/")
    log.Fatal(http.ListenAndServe(":8000", nil))
}