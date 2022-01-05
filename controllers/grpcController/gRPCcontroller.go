package grpcControllers

import (
	"context"
	"log"
	"uaAlert/pb"
	"uaAlert/repository"
	"uaAlert/services"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GrpcController struct {
	pb.UnimplementedClientsServiceServer
}

func (g *GrpcController) CreateStatus(ctx context.Context, req *pb.ClientsRequest) (*pb.ClientsResponse, error) {
	db := initmongodb()
	repo := repository.New(db)
	serv := services.New(repo)
	log.Printf("receive msg from %s", req.Client)
	cn := req.Client
	log.Println(cn)

	fn, err := repo.FindbyClientName(cn)
	if err != nil {
		log.Printf("grpcControllerCreateStatusFindbyClientError: %v", err)
	}

	status, err := serv.CheckStatus(cn, fn.LogFile)
	if err != nil {
		log.Printf("grpcControllerCreateStatusCheckStatusError: %v", err)
	}

	layout := "15:04:05"

	grpcClient := &pb.Client{
		Client:     cn,
		Status:     status.Status,
		Logtime:    status.LogTime.Format(layout),
		Systemtime: status.SystemTime.Format(layout),
		Difftime:   status.DiffTime.String(),
	}

	res := &pb.ClientsResponse{
		Status: grpcClient,
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
