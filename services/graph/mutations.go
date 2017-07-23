package graph

import (
	"../../models"
	"github.com/graphql-go/graphql"
)

// RootMutation is the start of the graphql node that allows mutations to db
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type:        UserType,
			Description: "Creates a new user",
			Args: graphql.FieldConfigArgument{
				"Email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				email, _ := params.Args["Email"].(string)

				user := models.User{Email: email}

				return user.SaveUser()
			},
		},
	},
})

// mutation {
//   createUser(email: "tim@gmail") {
//     email
//   }
// }
