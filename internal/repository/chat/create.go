package chat

import (
	"context"

	"github.com/Ghaarp/chat-server/internal/repository/chat/model"
	sq "github.com/Masterminds/squirrel"
)

func (repo *repo) Create(ctx context.Context, data *model.CreateRequest) (int64, error) {
	queryBuilder := sq.Insert("chats").PlaceholderFormat(sq.Dollar).
		Columns("author", "label").
		Values(data.Author, data.ChatName).
		Suffix("RETURNING id")

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	queryBuilderUsernames := sq.Insert("chat_users").PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "user_id")

	var chatid int64

	tx, err := repo.pool.Begin(ctx)
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	if err != nil {
		return 0, err
	}

	err = tx.QueryRow(ctx, query, args...).Scan(&chatid)
	if err != nil {
		return 0, err
	}

	for _, user := range data.Users {
		queryBuilderUsernames = queryBuilderUsernames.Values(chatid, user)
	}

	query, args, err = queryBuilderUsernames.ToSql()
	if err != nil {
		return 0, err
	}

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, err
	}

	return chatid, nil
}
