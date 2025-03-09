package main

import (
	"context"
	"fmt"
	"log"
	"net"

	generated "github.com/Ghaarp/chat-server/pkg/chat_v1"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	address = "localhost:"
	port    = 50060
)

type server struct {
	generated.UnimplementedChatV1Server
}

func (serv *server) Create(context context.Context, in *generated.CreateRequest) (*generated.CreateResponse, error) {
	log.Printf(color.GreenString("%v", in))

	return &generated.CreateResponse{}, nil
}

func (serv *server) Delete(context context.Context, in *generated.DeleteRequest) (*generated.DeleteResponse, error) {
	log.Printf(color.GreenString("%v", in))

	return &generated.DeleteResponse{}, nil
}

func (serv *server) SendMessage(context context.Context, in *generated.SendMessageRequest) (*generated.SendMessageResponse, error) {
	log.Printf(color.GreenString("%v", in))

	return &generated.SendMessageResponse{}, nil
}

func main() {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatal(err)
	}

	serverObj := grpc.NewServer()
	reflection.Register(serverObj)
	a := &server{}
	generated.RegisterChatV1Server(serverObj, a)

	log.Printf("Server started om %v", listener.Addr())

	if err := serverObj.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
