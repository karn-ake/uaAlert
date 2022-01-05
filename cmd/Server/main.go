package main

import (
	"log"
	"net"
	grpcservice "uaAlert/grpcService"
	"uaAlert/pb"

	"google.golang.org/grpc"
)

func main() {
	address := "127.0.0.1:8082"
	log.Printf("start server on %v", address)

	ctrl := grpcservice.GrpcServer{}
	// ctrlSvr := ctrl.CreateStatus()
	grpcServer := grpc.NewServer()

	pb.RegisterClientsServiceServer(grpcServer, &ctrl)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
