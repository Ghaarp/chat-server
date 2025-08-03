package main

import (
	"context"
	"log"

	"github.com/Ghaarp/chat-server/internal/app"
)

func main() {
	ctx := context.Background()
	application, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = application.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}

/*package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/Ghaarp/chat-server/internal/config"
	generated "github.com/Ghaarp/chat-server/pkg/chat_v1"
	sq "github.com/Masterminds/squirrel"
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
	pool *pgxpool.Pool
}

type request struct {
	query string
	args  []interface{}
}

func (serv *server) Create(context context.Context, in *generated.CreateRequest) (*generated.CreateResponse, error) {

	queryBuilder := sq.Insert("chats").PlaceholderFormat(sq.Dollar).
		Columns("author", "label").
		Values(in.Author, in.ChatName).
		Suffix("RETURNING id")

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return &generated.CreateResponse{}, err
	}

	queryBuilderUsernames := sq.Insert("chat_users").PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "user_id")

	var chatid int64

	tx, err := serv.pool.Begin(context)
	defer func() {
		if err != nil {
			tx.Rollback(context)
		}
	}()

	if err != nil {
		return &generated.CreateResponse{}, err
	}

	err = tx.QueryRow(context, query, args...).Scan(&chatid)
	if err != nil {
		return &generated.CreateResponse{}, err
	}

	for _, user := range in.Users {
		queryBuilderUsernames = queryBuilderUsernames.Values(chatid, user)
	}

	query, args, err = queryBuilderUsernames.ToSql()
	if err != nil {
		return &generated.CreateResponse{}, err
	}

	_, err = tx.Exec(context, query, args...)
	if err != nil {
		return &generated.CreateResponse{}, err
	}

	err = tx.Commit(context)
	if err != nil {
		return &generated.CreateResponse{}, err
	}

	return &generated.CreateResponse{Id: chatid}, nil
}

func (serv *server) Delete(context context.Context, in *generated.DeleteRequest) (*generated.DeleteResponse, error) {

	requests := make([]request, 3)

	queryChatBuilder := sq.Delete("chats").PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": in.Id})

	var err error
	requests[0].query, requests[0].args, err = queryChatBuilder.ToSql()
	if err != nil {
		return &generated.DeleteResponse{}, err
	}

	queryUsersBuilder := sq.Delete("chat_users").PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"chat_id": in.Id})

	requests[1].query, requests[1].args, err = queryUsersBuilder.ToSql()
	if err != nil {
		return &generated.DeleteResponse{}, err
	}

	queryMessagesBuilder := sq.Delete("messages").PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"chat_id": in.Id})

	requests[2].query, requests[2].args, err = queryMessagesBuilder.ToSql()
	if err != nil {
		return &generated.DeleteResponse{}, err
	}

	tx, err := serv.pool.Begin(context)
	if err != nil {
		return &generated.DeleteResponse{}, err
	}

	for _, request := range requests {
		_, err = tx.Exec(context, request.query, request.args...)
		if err != nil {
			return &generated.DeleteResponse{}, err
		}
	}

	err = tx.Commit(context)
	if err != nil {
		return &generated.DeleteResponse{}, err
	}

	return &generated.DeleteResponse{}, nil
}

func (serv *server) SendMessage(context context.Context, in *generated.SendMessageRequest) (*generated.SendMessageResponse, error) {
	builder := sq.Insert("messages").PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "author", "content").
		Values(in.Chatid, in.From, in.Text)

	query, args, err := builder.ToSql()
	if err != nil {
		return &generated.SendMessageResponse{}, err
	}

	_, err = serv.pool.Query(context, query, args...)
	if err != nil {
		return &generated.SendMessageResponse{}, err
	}

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

	turnOnServer(chatConfig, pool)

}

func turnOnServer(conf config.ChatConfig, pool *pgxpool.Pool) {
	listener, err := net.Listen("tcp", conf.Address())
	if err != nil {
		log.Fatal(err)
	}

	serverObj := grpc.NewServer()
	reflection.Register(serverObj)
	a := &server{}
	a.pool = pool
	generated.RegisterChatV1Server(serverObj, a)

	log.Printf("Server started on %v", listener.Addr())

	if err := serverObj.Serve(listener); err != nil {
		log.Fatal(err)
	}
}*/
