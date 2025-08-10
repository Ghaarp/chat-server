package chat

import (
	"context"

	"github.com/Ghaarp/chat-server/internal/repository/chat/model"
	sq "github.com/Masterminds/squirrel"
)

func (repo *repo) SendMessage(ctx context.Context, data *model.SendMessageRequest) error {

	builder := sq.Insert("messages").PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "author", "content").
		Values(data.Chatid, data.From, data.Text)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = repo.pool.Query(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
