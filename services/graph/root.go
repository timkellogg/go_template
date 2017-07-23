package graph

import (
	"log"

	"../../models"
	"github.com/graphql-go/graphql"
)

// RootQuery is the start of the graphql node that retrieves data
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type:        UserType,
			Description: "Returns a user",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				log.Println("RootMutation")
				// idQuery, isOK := params.Args["id"].(string)
				// if isOK {
				// retrieve user
				// }
				return models.User{}, nil
			},
		},
	},
})

// RootMutation is the start of the graphql node that allows mutations to db
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type:        UserType,
			Description: "Creates a new user",
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// email, _ := params.Args["email"].(string)
				// determine if valid

				// perform operation
				return models.User{}, nil
			},
		},
	},
})
