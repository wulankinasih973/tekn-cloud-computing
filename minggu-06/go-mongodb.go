package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Konfigurasi koneksi ke MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Membuat koneksi dengan database MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Mengatur timeout untuk koneksi
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Menguji koneksi ke server MongoDB
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Koneksi berhasil dibuat!")

	// Menutup koneksi ke MongoDB
	err = client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
