package main

import (
	"log"
	"net/http"

	"./services/graph"

	"github.com/gorilla/mux"
	"github.com/mnmtanish/go-graphiql"
)

// Route is a REST endpoint that maps common methods
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a collection of routes
type Routes []Route

var routes = Routes{
	Route{"Graphql", "GET", "/graphql", graph.Perform}, // figure out how to unify the GET and POST routes
	Route{"Graphql", "POST", "/graphql", graph.Perform},
	Route{"Graphiql", "GET", "/test", graphiql.ServeGraphiQL}, // get rid of in prod
}

// Cors - enable logging
type Cors struct {
	Log *log.Logger
}

// NewRouter establishes the root application router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
