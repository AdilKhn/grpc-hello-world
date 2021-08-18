package main

import (
	"fmt"
	"log"
	"net"

	"github.com/adilkhn/myrpc-server/chat"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))

	if err != nil {
		log.Fatal(err)
	}

	chatServer := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &chatServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
