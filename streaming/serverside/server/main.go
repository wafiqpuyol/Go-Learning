package main

import (
	"context"
	"fmt"
	"net"
	"server-streaming/notification"
	notificationProto "server-streaming/notification/proto/gen"
	myRedisClient "server-streaming/redis"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", notification.Address)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port:", notification.Address)

	grpcServer := grpc.NewServer()
	fmt.Println("GRPC Server started successfully")
	
	// Initialize Redis client
	ctx := context.Background()
	redisClient := myRedisClient.GetRedisClient(ctx)
	fmt.Println("Redis client initialized successfully", redisClient)

	// Register the notification service with the gRPC server
	notificationHandler := notification.NewNotificationHandler(redisClient)
	notificationProto.RegisterNotificationServiceServer(grpcServer, notificationHandler)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
