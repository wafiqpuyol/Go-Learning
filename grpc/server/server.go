package main

import (
	"context"
	"fmt"
	pb "grpc-server/proto/gen"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type calculateServer struct {
	pb.UnimplementedCalculateServer
}

func (s *calculateServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	log.Println("Received request:", req.A, req.B)
	result := req.A + req.B
	return &pb.AddResponse{Result: result}, nil
}

func main() {
	cert := "cert.pem"
	key := "key.pem"

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}

	// Set up a certificate and private key for the server
	credentials, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatalln("Failed to generate credentials", err)
	}

	// Create a gRPC server
	grpcServer := grpc.NewServer(grpc.Creds(credentials))
	log.Println("Server is running on port :50051")

	// Register the Calculate service with the server
	pb.RegisterCalculateServer(grpcServer, &calculateServer{})

	// Start the server & listen for requests
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("failed to serve:", err)
	}
}
