package main

import (
	"context"
	"log"
	"time"
	"uaAlert/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	address := "127.0.0.1:8082"
	log.Printf("dial server %s", address)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	grpcClient := pb.NewClientsServiceClient(conn)

	client := "BLP"
	req := &pb.ClientsRequest{
		Client: client,
	}
	log.Printf("sending %s request: %v", client, req)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := grpcClient.CreateStatus(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Print("laptop already exists")
		} else {
			log.Fatal("cannot create laptop: ", err)
		}
		return
	}

	log.Printf("gRPC response: %s", res.Status)
}
