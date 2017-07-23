package graph

import (
	"github.com/graphql-go/graphql"
)

// UserType - graphql user type
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"ID":    &graphql.Field{Type: graphql.Int},
		"Email": &graphql.Field{Type: graphql.String},
	},
})
