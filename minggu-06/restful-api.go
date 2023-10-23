package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoDBData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// Konfigurasi MongoDB
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = mongoClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect(ctx)

	router := gin.Default()

	router.GET("/mongodb-data", func(c *gin.Context) {
		collection := mongoClient.Database("db_mahasiswa").Collection("mahasiswa")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Perbaikan: tambahkan "cancel" di sini
		defer cancel() // Perbaikan: tambahkan "defer cancel()" di sini
		cur, err := collection.Find(ctx, bson.M{})
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

		c.JSON(http.StatusOK, data)
	})
	router.GET("/mysql-data", func(c *gin.Context) {
		// Implementasikan logika untuk mengambil data dari MySQL
		// dan mengembalikannya sebagai respons JSON
		var data []MySQLData
		// Implementasi pengambilan data dari MySQL di sini
		// Contoh:
		// db.Find(&data)
		c.JSON(http.StatusOK, data)
	})

	router.Run(":8080")
}
