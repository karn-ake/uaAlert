package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Client struct {
	LogFile    string `json:"LogFile" bson:"LogFile"`
	ClientName string `json:"ClientName" bson:"ClientName"`
}

type mongoRepository struct {
	db *mongo.Client
}

func New(db *mongo.Client) Repository {
	return &mongoRepository{db}
}

func (m *mongoRepository) FindAll() ([]Client, error) {
	if err := m.db.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to MongoDB alive")

	coll := m.db.Database("monitor").Collection("logfile")
	results, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	var clients []Client
	for results.Next(context.TODO()) {
		var client Client
		results.Decode(&client)
		clients = append(clients, client)
	}
	return clients, nil
}

func (m *mongoRepository) FindbyClientName(cn string) (*Client, error) {
	if err := m.db.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to MongoDB alive")

	coll := m.db.Database("monitor").Collection("logfile")
	result, err := coll.Find(context.TODO(), bson.D{{Key: "ClientName", Value: cn}})
	if err != nil {
		return nil, err
	}

	var client Client
	for result.Next(context.TODO()) {
		result.Decode(&client)
	}
	return &client, nil
}

func (m *mongoRepository) Update() error {
	docs, err := ConvJson()
	if err != nil {
		log.Fatalf("Can not load config %s", err)
	}
	coll := m.db.Database("monitor").Collection("logfile")
	var n int
	for _, doc := range docs {
		// log.Println(doc)
		a, _ := m.IsClientNameAdded(doc.ClientName)
		if a {
			log.Println("client have already been added")
		} else {
			result, _ := coll.InsertOne(context.TODO(), bson.D{{Key: "LogFile", Value: doc.LogFile}, {Key: "ClientName", Value: doc.ClientName}})
			n++
			log.Println(result.InsertedID)
		}
	}
	log.Printf("Insert %d rows", n)
	return nil
}

func (m *mongoRepository) DelAll() error {
	if err := m.db.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to MongoDB alive")

	coll := m.db.Database("monitor").Collection("logfile")
	result, err := coll.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}

func (m *mongoRepository) IsClientNameAdded(cn string) (bool, error) {
	var client *Client
	client, err := m.FindbyClientName(cn)
	if err != nil {
		log.Fatal(err)
	}

	if cn == client.ClientName {
		return true, nil
	}

	return false, nil
}
