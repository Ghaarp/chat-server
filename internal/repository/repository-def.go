package repository

import (
	"context"

	"github.com/Ghaarp/chat-server/internal/repository/chat/model"
)

type ChatRepository interface {
	Create(ctx context.Context, data *model.CreateRequest) (int64, error)
	Delete(ctx context.Context, id int64) error
	SendMessage(ctx context.Context, data *model.SendMessageRequest) error
	ClosePool(ctx context.Context)
}
