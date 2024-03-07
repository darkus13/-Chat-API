package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/darkus13/-Chat_API/pkg/chat_v1"
)

const (
	dbDSN     = "host=localhost port=5433 dbname=chat user=darkus password=andrej sslmode=disable"
	grpcPort  = 50051
	chatId    = "id"
	chatName  = "name"
	chatText  = "text"
	tableChat = "chat"
)

type server struct {
	desc.UnimplementedChatV1Server
	db *pgxpool.Pool
	qb sq.StatementBuilderType
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	name := req.GetUsername()

	InsertBuilder := s.qb.Insert(tableChat).
		PlaceholderFormat(sq.Dollar).
		Columns(chatName).
		Values(name).
		Suffix("RETURNING id")

	query, args, err := InsertBuilder.ToSql()
	if err != nil {
		_ = fmt.Errorf("failed to build query: %v", err)
	}

	var chatID int64

	err = s.db.QueryRow(ctx, query, args...).Scan(&chatID)
	if err != nil {
		_ = fmt.Errorf("failed to insert user: %v", err)
	}

	log.Printf("inserted user with ID: %d", chatID)

	return &desc.CreateResponse{
		Id: chatID,
	}, nil

}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	ID := req.GetId()

	DeleteBuilder := s.qb.Delete(tableChat).
		Where(sq.Eq{chatId: ID})

	query, args, err := DeleteBuilder.ToSql()
	if err != nil {
		_ = fmt.Errorf("failed to build query: %v", err)
	}

	row, err := s.db.Exec(ctx, query, args...)
	if err != nil {
		_ = fmt.Errorf("failed to delete user: %v", err)
	}

	log.Printf("delete %d rows", row.RowsAffected())

	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {

	text := req.GetText()

	SelectBuilder := s.qb.Insert(chatText).
		Columns(chatText).
		Values(text).
		Prefix("RETURNING text")

	query, args, err := SelectBuilder.ToSql()
	if err != nil {
		_ = fmt.Errorf("failed to build query: %v", err)
	}

	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		_ = fmt.Errorf("failed to send massage db : %v", err)
	}

	return &emptypb.Empty{}, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	pgxConfig, err := pgxpool.ParseConfig(dbDSN)
	if err != nil {
		_ = fmt.Errorf("failed to patde config: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		_ = fmt.Errorf("failed to connect to postgres: %v", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		_ = fmt.Errorf("ping to postgres failed: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		_ = fmt.Errorf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{
		db: pool,
		qb: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
