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

	// Create a Calculate client
	calculateClient := mainpb.NewCalculateClient(connection)

	// Call the Add rpc
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := calculateClient.Add(ctx, &mainpb.AddRequest{A: 40, B: 64})
	if err != nil {
		log.Fatalf("Error while calling Add RPC: %v", err)
	}

	log.Println("Result:", res.Result)
}
