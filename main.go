package main

import (
	"context"
	"log"
	"net"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"

	pb "chat/proto/chat"
)

type chatServer struct {
	pb.UnimplementedChatServiceServer
	redisClient *redis.Client
}

func (s *chatServer) AddMessage(ctx context.Context, req *pb.AddMessageRequest) (*pb.AddMessageResponse, error) {
	// Добавление сообщения в чат с использованием Redis
	err := s.redisClient.RPush(ctx, req.ChatId, req.Message).Err()
	if err != nil {
		return nil, err
	}

	return &pb.AddMessageResponse{}, nil
}

func (s *chatServer) GetMessages(ctx context.Context, req *pb.GetMessagesRequest) (*pb.GetMessagesResponse, error) {
	// Получение списка сообщений из чата с использованием Redis
	messages, err := s.redisClient.LRange(ctx, req.ChatId, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	response := &pb.GetMessagesResponse{
		Messages: messages,
	}

	return response, nil
}

func (s *chatServer) GetChats(ctx context.Context, req *pb.GetChatsRequest) (*pb.GetChatsResponse, error) {
	// Получение списка всех чатов с использованием Redis
	keys, err := s.redisClient.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	response := &pb.GetChatsResponse{
		Chats: keys,
	}

	return response, nil
}

func main() {
	// Создание подключения к Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Проверка подключения к Redis
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	// Создание gRPC сервера
	server := grpc.NewServer()

	// Регистрация chatServer в качестве обработчика gRPC запросов
	pb.RegisterChatServiceServer(server, &chatServer{
		redisClient: redisClient,
	})

	// Указание адреса и порта для прослушивания соединений
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Запуск gRPC сервера
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
