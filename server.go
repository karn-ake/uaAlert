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
	// fn, _ := repo.FindbyClientName(client)
	// // at, _ := serv.GetAllTimes(fn.LogFile)
	// cs, _ := serv.CheckStatus(client, fn.LogFile)
	// log.Printf("\nSystem time: %v\nLog time: %v\nDifferent time: %v",at.SystemTime,at.LogTime,at.DiffTime)
	// log.Println(*cs)
	// bs, _ := serv.RevFile(fn.LogFile)

	// for _, b := range *bs{
	// 	log.Println(b)
	// }

	// result, err := repo.DelAll()
	// if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(result)

	// repo.Update()
	// clients, err := repo.FindAll()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var cl repository.Client
	// for _, cl = range clients {
	// 	log.Println(cl.ClientName,cl.LogFile)
	// }

	// for _, cl = range client {
	// 	log.Printf("Client name: %s, Log file on: %s", cl.ClientName, cl.LogFile)
	// }

	// client, err := repo.FindbyClientName("BLP")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("Client name: %s, Log file on: %s", client.ClientName, client.LogFile)

	// if b, err := repo.IsClientNameAdded(" BLP"); err == nil {
	// 	if b {
	// 		log.Println("This client name is already added")
	// 	} else {
	// 		log.Println("This client name is not added")
	// 	}
	// }
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
