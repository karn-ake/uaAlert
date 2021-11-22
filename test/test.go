package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Client struct {
	LogFile    string `json:"FileName" bson:"FileName"`
	ClientName string `json:"ClientName" bson:"ClientName"`
}

func main() {
	DelAll()
	Update()
}

func ConvJson() (clients []Client, err error) {
	file, err := os.Open("config.txt")
	if err != nil {
		log.Fatalf("Can't open file %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// var clients []Client
	for scanner.Scan() {
		s := scanner.Text()
		sp := strings.Split(s, ",")
		con := Client{
			LogFile:    sp[0],
			ClientName: sp[1],
		}
		// fmt.Println(sp[0])
		clients = append(clients, con)
	}

	// j, err := json.Marshal(clients)
	// if err != nil {
	// 	log.Fatalf("cannot marshal: %v", err)
	// }
	return clients, nil
}

func Update() error {
	docs, _ := ConvJson()
	db := initmongodb()
	defer db.Disconnect(context.TODO())
	coll := db.Database("monitor").Collection("logfile")
	var n int
	for _, doc := range docs {
		// log.Println(doc)
		result, _ := coll.InsertOne(context.TODO(), bson.D{{Key: "LogFile", Value: doc.LogFile}, {Key: "ClientName", Value: doc.ClientName}})
		n++
		log.Println(result.InsertedID)
	}
	log.Printf("Insert %d rows", n)
	return nil
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

func DelAll() (*mongo.DeleteResult, error) {
	db := initmongodb()
	defer db.Disconnect(context.TODO())

	if err := db.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to MongoDB alive")

	coll := db.Database("monitor").Collection("logfile")
	result, err := coll.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	return result, nil
}
