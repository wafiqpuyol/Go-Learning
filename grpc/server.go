package main

import (
	"context"
	"fmt"
	pb "grpc-server/proto/gen"
	"log"
	"net"

	"google.golang.org/grpc"
)

type calculateServer struct {
	pb.UnimplementedCalculateServer
}

func (calculateServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	result := req.A + req.B
	return &pb.AddResponse{Result: result}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}

	grpcServer := grpc.NewServer()
	log.Println("Server is running on port :50051")

	pb.RegisterCalculateServer(grpcServer, &calculateServer{})

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("failed to serve:", err)
	}
}
