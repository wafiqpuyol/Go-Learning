package notification

import (
	"fmt"
	notificationProto "server-streaming/notification/proto/gen"

	"github.com/redis/go-redis/v9"
)

type handler struct {
	redisClient *redis.Client
	notificationProto.UnimplementedNotificationServiceServer
}

func NewNotificationHandler(redisClient *redis.Client) *handler {
	return &handler{
		redisClient: redisClient,
	}
}

func GetNotification(req *notificationProto.NotificationRequest, stream notificationProto.NotificationService_GetNotificationServer) error {
	userId := req.GetUserId()
	fmt.Println("User ID:", userId)
	return nil
}
