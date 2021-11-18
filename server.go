package main

import (
	"context"
	"log"
	"uaAlert/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	db := initmongodb()

	repo := repository.New(db)
	// repo.Update()

	// result, err := repo.DelAll()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(result)

	clients, err := repo.FindAll()
	if err != nil {
		log.Fatal(err)
	}

	var cl repository.Client
	for _, cl = range clients {
		log.Println(cl.ClientName)
	}

	client, err := repo.FindbyClientName("BLP")
	if err != nil {
		log.Fatal(err)
	}

	for _, cl = range client {
		log.Printf("Client name: %s, Log file on: %s", cl.ClientName, cl.LogFile)
	}
}

func initmongodb() *mongo.Client {
	const uri = "mongodb://root:123456@192.168.170.131:27017/?maxPoolSize=20&w=majority"
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected")
	return client
}
