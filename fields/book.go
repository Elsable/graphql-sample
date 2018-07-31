package fields

import (
	"errors"
	"log"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/stobita/graphql-sample/service"
)

var BookField = &graphql.Field{
	Type:        bookType,
	Description: "Get Single Book",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		stringBookID, ok := params.Args["id"].(string)
		bookID, _ := strconv.ParseInt(stringBookID, 10, 64)
		log.Println(bookID)
		if ok {
			book := service.NewBook()
			if book.FindByID(bookID) {
				return book, nil
			}
		}
		return nil, errors.New("invalid bookID")
	},
}

var BooksField = &graphql.Field{
	Type:        graphql.NewList(bookType),
	Description: "Get Multiple Book",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return service.NewBook().GetAll(), nil
	},
}

var bookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})
