package graph

import (
	"log"

	"github.com/graphql-go/graphql"
)

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    RootQuery,
	Mutation: RootMutation,
})

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		// handle errors better
		log.Printf("Failed executing graphql query. Errors: %v", result.Errors)
	}
	return result
}

// Process the api to the graph package. Takes a http request object and calls the graph layer
// to perform an operation and return the results
// func Process(query string) {

// }
