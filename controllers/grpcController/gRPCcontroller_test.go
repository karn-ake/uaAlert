package grpcControllers_test

import (
	"context"
	"log"
	"testing"
	"uaAlert/pb"
	"uaAlert/repository"
	"uaAlert/services"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCreateStatus(t *testing.T) {
	t.Parallel()
	cn := "BLP"
	db := initmongodb()
	repo := repository.New(db)
	serv := services.New(repo)
	// tcs := controllers.NewGrpcController(repo, serv)
	fn, _ := repo.FindbyClientName(cn)
	tc, _ := serv.CheckStatus(cn, fn.LogFile)
	// tc := tcs.CreateStatus(context.TODO())
	layout := "15:04:05"

	grpc := &pb.Client{
		Client:     tc.Client,
		Status:     tc.Status,
		Logtime:    tc.LogTime.Format(layout),
		Systemtime: tc.SystemTime.Format(layout),
		Difftime:   tc.DiffTime.String(),
	}
	t.Log(grpc)
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
