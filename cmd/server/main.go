package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/Ghaarp/chat-server/internal/config"
	generated "github.com/Ghaarp/chat-server/pkg/chat_v1"
	"github.com/fatih/color"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

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

	flag.Parse()
	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Print("Unable to load .env")
	}

	chatConfig, err := config.NewChatConfig()
	if err != nil {
		log.Fatal("Unable to load chat config")
	}

	dbconfig, err := config.NewDBConfig()
	if err != nil {
		log.Fatal("Unable to load DB config")
	}

	pool, err := pgxpool.Connect(ctx, dbconfig.DSN())
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	turnOnServer(chatConfig)

}

func turnOnServer(conf config.ChatConfig) {
	listener, err := net.Listen("tcp", conf.Address())
	if err != nil {
		log.Fatal(err)
	}

	serverObj := grpc.NewServer()
	reflection.Register(serverObj)
	a := &server{}
	generated.RegisterChatV1Server(serverObj, a)

	log.Printf("Server started on %v", listener.Addr())

	if err := serverObj.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
