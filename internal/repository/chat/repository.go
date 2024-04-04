package chat

import (
	"github.com/darkus13/-Chat_API/internal/client/db"
	"github.com/darkus13/-Chat_API/internal/repository"
)

const (
	returnID = "RETURNING id"
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
