package graph

import (
	"log"
	"net/http"

	"encoding/json"

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

// Perform a graph query and return the results
func Perform(w http.ResponseWriter, r *http.Request) {
	result := executeQuery(r.URL.Query().Get("query"), schema)
	json.NewEncoder(w).Encode(result)
}
