package chat

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/darkus13/-Chat_API/internal/client/db"
	"github.com/darkus13/-Chat_API/internal/repository/chat/model"
)

func (r *repo) SendMessage(ctx context.Context, info *model.Info) error {
	builderInsert := sq.Insert(chat).
		Columns(chatID, userID, text).
		Values(info.Users, info.ChatID, info.Content)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.SendMessage",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
