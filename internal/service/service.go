package service

import (
	"context"

	"github.com/darkus13/-Chat_API/internal/repository/chat/model"
)

type ChatService interface {
	Create(ctx context.Context, info *model.Info) (int64, error)
	Delete(ctx context.Context, info *model.Info) error
	SendMessage(ctx context.Context, info *model.Info) error
}
