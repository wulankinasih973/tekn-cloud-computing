package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var mongoClient *mongo.Client

func initMongoDB() {
	var err error
	mongoClient, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = mongoClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

var mongoType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "MongoDBData",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"mongoData": &graphql.Field{
				Type:        graphql.NewList(mongoType),
				Description: "Mengambil data dari MongoDB",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					collection := mongoClient.Database("db_mahasiswa").Collection("mahasiswa")
					ctx := context.TODO()
					cur, err := collection.Find(ctx, nil)
					if err != nil {
						log.Fatal(err)
					}
					defer cur.Close(ctx)

					var data []MongoDBData
					for cur.Next(ctx) {
						var result MongoDBData
						err := cur.Decode(&result)
						if err != nil {
							log.Fatal(err)
						}
						data = append(data, result)
					}
					return data, nil
				},
			},
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: rootQuery,
	},
)

func main() {
	initMongoDB()

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)

	fmt.Println("Server GraphQL untuk MongoDB berjalan pada :8080/graphql")
	http.ListenAndServe(":8080", nil)
}
