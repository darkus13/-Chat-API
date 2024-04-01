package chat

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"

	"github.com/darkus13/-Chat_API/internal/client/db"
	"github.com/darkus13/-Chat_API/internal/repository/chat/model"
)

func (r *repo) Delete(ctx context.Context, info *model.Info) error {
	builderDelete := sq.Delete(chat).
		Where(sq.Eq{userID: userID})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "chat.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		log.Printf("failed to query: %v", err)
	}

	return nil

}
