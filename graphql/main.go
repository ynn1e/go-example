package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
)

func main() {
	http.HandleFunc("/graphql", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("listen and serve error: %s", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Hello, GraphQL!", nil
			},
		},
		"now": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return time.Now().Format(time.DateTime), nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Printf("Error creating schema: %s", err)

		return
	}

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: r.URL.Query().Get("query"),
	})
	if len(result.Errors) > 0 {
		log.Printf("Errors: %s", result.Errors)

		return
	}
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Printf("encode error: %s", err)

		return
	}
}
