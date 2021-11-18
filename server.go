package main

import (
	"context"
	"fmt"
	"log"
	"uaAlert/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	db := initmongodb()
	_ = db

	repo := repository.New(db)
	// repo.Update()

	result, err := repo.DelAll()
	if err != nil {log.Fatal(err)}
	fmt.Println(result)
	
	clients, err := repo.FindAll()
	if err != nil {
		log.Fatal(err)
	}

	var cl repository.Client
	for _ , cl = range clients {
		fmt.Println(cl.ClientName)
	}

	client, err := repo.FindbyClientName("BLP")
	if err != nil { log.Fatal(err)}

	// var cl repository.Client
	for _ , cl = range client {
		fmt.Println(cl.ClientName)
	}
}

func initmongodb() *mongo.Client {
	const uri = "mongodb://root:123456@192.168.170.131:27017/?maxPoolSize=20&w=majority"
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected")
	return client
}
