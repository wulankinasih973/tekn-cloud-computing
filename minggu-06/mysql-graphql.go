package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLData struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

var db *gorm.DB

func initMySQLDB() {
	dsn := "wulan:wulwul21@tcp(localhost:3306)/db_mahasiswa"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
}

var mysqlType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "MySQLData",
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
			"mysqlData": &graphql.Field{
				Type:        graphql.NewList(mysqlType),
				Description: "Mengambil data dari MySQL",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					var data []MySQLData
					db.Find(&data)
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
	initMySQLDB()

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)

	fmt.Println("Server GraphQL untuk MySQL berjalan pada :8080/graphql")
	http.ListenAndServe(":8080", nil)
}
