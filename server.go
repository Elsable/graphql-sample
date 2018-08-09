package main

import (
	"bytes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/stobita/graphql-sample/fields"
	"github.com/stobita/graphql-sample/middlewares"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware)
	r.POST("/graphql", getBook)
	port := os.Getenv("PORT")
	r.Run(":" + port)
}

func getBook(c *gin.Context) {
	bufferBody := new(bytes.Buffer)
	bufferBody.ReadFrom(c.Request.Body)
	body := bufferBody.String()
	// requestBody, _ := ioutil.ReadAll(c.Request.Body)
	// body := fmt.Sprintf("%s", body)
	log.Println(body)
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
		"user":  fields.UserField,
		"users": fields.UsersField,
		"book":  fields.BookField,
		"books": fields.BooksField,
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return
	}
	params := graphql.Params{Schema: schema, RequestString: body}
	result := graphql.Do(params)
	c.JSON(200, result)
	return
}
