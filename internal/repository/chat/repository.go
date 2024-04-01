package chat

import (
	"github.com/darkus13/-Chat_API/internal/client/db"
	"github.com/darkus13/-Chat_API/internal/repository"
)

const (
	dbDSN    = "host=localhost port=5433 dbname=chat user=darkus password=andrej sslmode=disable"
	grpcPort = 50051
	chat     = "chat"
	chatID   = "chat_id"
	userID   = "user_id"
	text     = "text"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{
		db: db,
	}
}
