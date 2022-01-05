package grpcservice

import (
	"context"
	"log"
	"uaAlert/pb"
	"uaAlert/repository"
	"uaAlert/services"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GrpcServer struct {
	pb.UnimplementedClientsServiceServer
}

func (*GrpcServer) CreateStatus(ctx context.Context, req *pb.ClientsRequest) (*pb.ClientsResponse, error) {
	db := initmongodb()
	repo := repository.New(db)
	serv := services.New(repo)

	cn := req.GetClient()
	fn, _ := repo.FindbyClientName(cn)
	client, _ := serv.CheckStatus(cn, fn.LogFile)

	layout := "15:04:05"

	cs := &pb.Client{
		Client:     cn,
		Status:     client.Status,
		Logtime:    client.LogTime.Format(layout),
		Systemtime: client.SystemTime.Format(layout),
		Difftime:   client.DiffTime.String(),
	}

	res := &pb.ClientsResponse{
		Status: cs,
	}

	return res, nil
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
