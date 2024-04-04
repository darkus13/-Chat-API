package chat

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"

	"github.com/darkus13/-Chat_API/internal/client/db"
	"github.com/darkus13/-Chat_API/internal/repository/chat/model"
)

func (r *repo) Create(ctx context.Context, info *model.Info) (int64, error) {
	builderInsert := sq.Insert(chat).
		Columns(chatID).
		Values(info.ChatID).
		Suffix(returnID)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return 0, err
	}

	q := db.Query{
		Name:     "chat.Create",
		QueryRaw: query,
	}

	var chatiD int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&chatiD)
	if err != nil {
		return 0, err
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		log.Fatalf("failed to added data in db: %v", err)
		return 0, err
	}

	return chatiD, nil
}
