package main

import (
	"context"
	"fmt"
	mainpb "grpc-client/proto/gen"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// create certificate
	cert := "cert.pem"
	credentials, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		log.Fatalln("Failed to load certificates", err)
	}

	// Make a connection to the server
	connection, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(credentials))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer connection.Close()
	fmt.Println("Connected to server")

	// Create a RPC clients
	calculateClient := mainpb.NewCalculateClient(connection)
	greeterClient := mainpb.NewGreeterClient(connection)

	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call RPCs
	res, err := calculateClient.Add(ctx, &mainpb.AddRequest{A: 40, B: 64})
	if err != nil {
		log.Fatalf("Error while calling Add RPC: %v", err)
	}
	log.Println("Result from Add():", res.Result)

	res1, err := greeterClient.SayHello(ctx, &mainpb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling  RPC: %v", err)
	}
	log.Println("Result from SayHello():", res1.Greet)

}
