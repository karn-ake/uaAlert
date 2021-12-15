package main

import (
	"context"
	"log"
	"uaAlert/controllers"
	"uaAlert/repository"
	"uaAlert/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	db := initmongodb()
	repo := repository.New(db)
	serv := services.New(repo)
	cont := controllers.NewFiberController(repo, serv)

	//repo.DelAll()
	port := viper.GetString("app.port")
	log.Println(port)

	// Fiber routeconfiguration
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Fiber up and running")
	})
	app.Get("/api/updateconfig", cont.UpdateConfig)
	app.Get("/api/clearconfig", cont.DeleteConfig)
	app.Get("/api/:client", cont.ClientController)
	log.Fatal(app.Listen(port))
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

func init() {
	viper.SetConfigName("config")
	// viper.SetConfigFile("yaml")
	viper.AddConfigPath("D:\\Go\\src\\uaAlert\\")

	//Config path on OMS1-UAT
	//viper.AddConfigPath("D:\\Scripts\\Go\\uaAlert\\")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalln("config file is not found")
		} else {
			log.Fatalf("cannot load config file: %v", err)
		}
	}
	log.Println("config file was loaded")
}
