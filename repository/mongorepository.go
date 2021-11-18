package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Client struct {
	LogFile    string `json:"logfile" bson:"logfile"`
	ClientName string `json:"clientname" bson:"clientname"`
}

type mongoRepository struct {
	db *mongo.Client
}

func New(db *mongo.Client) Repository {
	return mongoRepository{db}
}

func (m mongoRepository) FindAll() ([]Client, error) {
	if err := m.db.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Connection to MongoDB alive")

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

func (m mongoRepository) FindbyClientName(cn string) ([]Client, error) {
	if err := m.db.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Connection to MongoDB alive")

	coll := m.db.Database("monitor").Collection("logfile")
	results, err := coll.Find(context.TODO(), bson.D{{Key: "clientName", Value: cn}})
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

func (m mongoRepository) Update() error {
	docs := []interface{}{
		bson.D{{Key: "LogFile", Value: "D:\\N2N\\UA_NOMURA\\BS_BLP\\logs\\log.txt"}, {Key: "ClientName", Value: "BLP"}},
		bson.D{{Key: "LogFile", Value: "D:\\N2N\\UA_NOMURA\\BS_CLV\\logs\\log.txt"}, {Key: "ClientName", Value: "KASIKORN"}},
		bson.D{{Key: "LogFile", Value: "D:\\N2N\\UA_NOMURA\\BS_CLV\\logs\\log.txt"}, {Key: "ClientName", Value: "KSAMCRD"}},
		bson.D{{Key: "LogFile", Value: "D:\\N2N\\UA_NOMURA\\BS_CLV\\logs\\log.txt"}, {Key: "ClientName", Value: "KTAMCRD"}},
		bson.D{{Key: "LogFile", Value: "D:\\N2N\\UA_NOMURA\\BS_CLV\\logs\\log.txt"}, {Key: "ClientName", Value: "MFCAMCRD"}},
		bson.D{{Key: "LogFile", Value: "D:\\N2N\\UA_NOMURA\\BS_CLV\\logs\\log.txt"}, {Key: "ClientName", Value: "SCBAMCRD"}},
		bson.D{{Key: "LogFile", Value: "D:\\N2N\\UA_NOMURA\\BS_INS\\logs\\log.txt"}, {Key: "ClientName", Value: "INSTINET"}},
		bson.D{{Key: "LogFile", Value: "D:\\N2N\\UA_NOMURA\\BS_ALDN\\logs\\log.txt"}, {Key: "ClientName", Value: "NYFIX"}},
	}
	coll := m.db.Database("monitor").Collection("logfile")
	result, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		return err
	}

	fmt.Printf("%v", len(result.InsertedIDs))
	return nil
}

func (m mongoRepository) DelAll() (*mongo.DeleteResult, error) {
	if err := m.db.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Connection to MongoDB alive")

	coll := m.db.Database("monitor").Collection("logfile")
	result, err := coll.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	return result, nil
}
