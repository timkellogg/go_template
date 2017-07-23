package graph

import (
	"encoding/json"
	"net/http"

	"log"

	"github.com/graphql-go/graphql"
	graphiql "github.com/mnmtanish/go-graphiql"
)

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    RootQuery,
	Mutation: RootMutation,
})

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	log.Println(query)
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	return result
}

func sendError(w http.ResponseWriter, err error) {
	log.Println(err)
	w.WriteHeader(500)
	w.Write([]byte(err.Error()))
}

// Perform a graph query and return the results
func Perform(w http.ResponseWriter, r *http.Request) {
	req := &graphiql.Request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		sendError(w, err)
		return
	}

	result := executeQuery(req.Query, schema)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		sendError(w, err)
	}

	return
}
