package graph

import (
	"errors"
	"strconv"

	"../../config"
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
				"ID": &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: func(params graphql.ResolveParams) (models.User, error) {
				id, err := strconv.Atoi(params.Args["ID"].(string))

				var user models.User

				if err == nil {

					record := config.DB.Find(&user, id)

					return record, err
				}

				return models.User{}, errors.New("Not Found")
			},
		},
	},
})
