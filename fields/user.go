package fields

import (
	"errors"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/stobita/graphql-sample/service"
)

var UserField = &graphql.Field{
	Type:        userType,
	Description: "Get Single User",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		stringUserID, ok := params.Args["id"].(string)
		userID, _ := strconv.ParseInt(stringUserID, 10, 64)
		if ok {
			user := service.NewUser()
			if user.FindByID(userID) {
				return user, nil
			}
		}
		return nil, errors.New("invalid userID")
	},
}

var UsersField = &graphql.Field{
	Type:        graphql.NewList(userType),
	Description: "Get Multiple User",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return service.NewUser().GetAll(), nil
	},
}

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})
