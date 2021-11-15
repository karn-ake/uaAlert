package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

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
