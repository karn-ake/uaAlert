package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Client struct {
	logFile string
	clientName string
}

type mongoRepository struct {
	db *mongo.Client
}

func New(db *mongo.Client) Repository {
	return mongoRepository{db}
}

func (m mongoRepository) FindAll() ([]string, error) {
	if err := m.db.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Connection to MongoDB alive")

	return nil, nil
}

func (m mongoRepository) Update() {
	docs := []interface{}{
		bson.D{{"logFile", "My Brilliant Friend"}, {"clientName", "Elena Ferrante"}},
		bson.D{{"logFile", "Lucy"}, {"clientName", "Jamaica Kincaid"}},
		bson.D{{"logFile", "Cat's Cradle"}, {"clientName", "Kurt Vonnegut Jr."}},
	}
	coll := m.db.Database("monitor").Collection("log")
	result, err := coll.InsertMany(context.TODO(),docs)
	if err != nil {
		fmt.Println(err)
	}

}
