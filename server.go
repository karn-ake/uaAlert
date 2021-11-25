package main

import (
	"context"
	"log"
	"uaAlert/controllers"
	"uaAlert/repository"
	"uaAlert/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	db := initmongodb()
	repo := repository.New(db)
	serv := services.New(repo)
	cont := controllers.NewFiberController(repo, serv)
	// rmux := routes.New(cont)
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Fiber up and running")
	})

	app.Get("/api/:client", cont.ClientController)

	const port string = ":8082"
	log.Fatal(app.Listen(port))
	// rmux.GET("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Up and Running")
	// })
	// rmux.GET("/api/{client}", cont.ClientController)
	// rmux.SERV(port)
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
