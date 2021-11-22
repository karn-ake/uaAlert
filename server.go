package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"uaAlert/controllers"
	"uaAlert/repository"
	"uaAlert/routes"
	"uaAlert/services"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	db := initmongodb()
	repo := repository.New(db)
	serv := services.New(repo)
	cont := controllers.New(serv, repo)
	rmux := routes.New(cont)

	const port string = ":8082"
	rmux.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and Running")
	})
	rmux.GET("/api/{client}", cont.ClientController)
	rmux.SERV(port)
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
