package graph

import (
	"github.com/graphql-go/graphql"
)

// UserType - graphql user type
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{ Type: graphql.String },
		"email": &graphql.Field{ Type: graphql.String },
	}
})
