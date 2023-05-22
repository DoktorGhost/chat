package main

import (
	"context"
	"testing"

	pb "chat/proto/chat"

	"github.com/redis/go-redis/v9"
)

func TestAddMessage(t *testing.T) {
	// Создаем экземпляр Redis-клиента для тестов
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	// Создаем экземпляр chatServer с использованием Redis-клиента для тестов
	server := &chatServer{
		redisClient: redisClient,
	}

	// Создаем фиктивный контекст
	ctx := context.Background()

	// Создаем запрос на добавление сообщения
	req := &pb.AddMessageRequest{
		ChatId:  "chat1",
		Message: "Hello, world!",
	}

	// Вызываем метод AddMessage и проверяем ошибку
	_, err := server.AddMessage(ctx, req)
	if err != nil {
		t.Errorf("AddMessage failed: %v", err)
	}

	// Проверяем, что сообщение успешно добавлено в Redis
	messages, err := redisClient.LRange(ctx, "chat1", 0, -1).Result()
	if err != nil {
		t.Errorf("Failed to get messages from Redis: %v", err)
	}
	if len(messages) != 1 || messages[0] != "Hello, world!" {
		t.Errorf("Unexpected messages in Redis: %v", messages)
	}
}

func TestGetMessages(t *testing.T) {
	// Создаем экземпляр Redis-клиента для тестов
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	// Создаем экземпляр chatServer с использованием Redis-клиента для тестов
	server := &chatServer{
		redisClient: redisClient,
	}

	// Создаем фиктивный контекст
	ctx := context.Background()

	// Помещаем сообщения в Redis
	err := redisClient.RPush(ctx, "chat1", "Message 1", "Message 2").Err()
	if err != nil {
		t.Fatalf("Failed to push messages to Redis: %v", err)
	}

	// Создаем запрос на получение сообщений
	req := &pb.GetMessagesRequest{
		ChatId: "chat1",
	}

	// Вызываем метод GetMessages и проверяем результат
	res, err := server.GetMessages(ctx, req)
	if err != nil {
		t.Errorf("GetMessages failed: %v", err)
	}

	expectedMessages := []string{"Message 1", "Message 2"}
	if len(res.Messages) != len(expectedMessages) {
		t.Errorf("Unexpected number of messages. Expected: %d, Got: %d", len(expectedMessages), len(res.Messages))
	}

	for i, msg := range res.Messages {
		if msg != expectedMessages[i] {
			t.Errorf("Unexpected message at index %d. Expected: %s, Got: %s", i, expectedMessages[i], msg)
		}
	}
}

func TestGetChats(t *testing.T) {
	// Создаем экземпляр Redis-клиента для тестов
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	// Создаем экземпляр chatServer с использованием Redis-клиента для тестов
	server := &chatServer{
		redisClient: redisClient,
	}

	// Создаем фиктивный контекст
	ctx := context.Background()

	// Помещаем ключи чатов в Redis
	err := redisClient.Set(ctx, "chat1", "", 0).Err()
	if err != nil {
		t.Fatalf("Failed to set chat keys in Redis: %v", err)
	}
	err = redisClient.Set(ctx, "chat2", "", 0).Err()
	if err != nil {
		t.Fatalf("Failed to set chat keys in Redis: %v", err)
	}

	// Создаем запрос на получение чатов
	req := &pb.GetChatsRequest{}

	// Вызываем метод GetChats и проверяем результат
	res, err := server.GetChats(ctx, req)
	if err != nil {
		t.Errorf("GetChats failed: %v", err)
	}

	expectedChats := []string{"chat1", "chat2"}
	if len(res.Chats) != len(expectedChats) {
		t.Errorf("Unexpected number of chats. Expected: %d, Got: %d", len(expectedChats), len(res.Chats))
	}

	for i, chat := range res.Chats {
		if chat != expectedChats[i] {
			t.Errorf("Unexpected chat at index %d. Expected: %s, Got: %s", i, expectedChats[i], chat)
		}
	}
}
