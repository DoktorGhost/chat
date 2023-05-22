package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"

	pb "chat/proto/chat"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)

	var com, chatId string
	for {
		fmt.Println("M - addMessage, GM - getMessage, GC - getChat")

		fmt.Scanln(&com)
		switch com {
		case "M":
			fmt.Println("chatId")
			fmt.Scan(&chatId)
			fmt.Println("message")
			read := bufio.NewReader(os.Stdin)
			mess, err := read.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}

			addMessageReq := &pb.AddMessageRequest{
				ChatId:  chatId,
				Message: mess,
			}
			_, err = client.AddMessage(context.Background(), addMessageReq)
			if err != nil {
				log.Fatalf("Failed to add message: %v", err)
			}

		case "GM":
			fmt.Println("chatId:")
			fmt.Scan(&chatId)
			getMessagesReq := &pb.GetMessagesRequest{
				ChatId: chatId,
			}
			getMessagesRes, err := client.GetMessages(context.Background(), getMessagesReq)
			if err != nil {
				log.Fatalf("Failed to get messages: %v", err)
			}
			for _, message := range getMessagesRes.Messages {
				log.Println(message)
			}
		case "GC":
			getChatsReq := &pb.GetChatsRequest{}
			getChatsRes, err := client.GetChats(context.Background(), getChatsReq)
			if err != nil {
				log.Fatalf("Failed to get chats: %v", err)
			}

			for i, chat := range getChatsRes.Chats {
				fmt.Printf("%d) %s; ", i+1, chat)
			}
			fmt.Println()
		}
	}

}
