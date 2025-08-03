package service

import (
	"context"

	"github.com/Ghaarp/chat-server/internal/service/chat/model"
)

type ChatService interface {
	Create(ctx context.Context, data *model.CreateRequest) (int64, error)
	Delete(ctx context.Context, id int64) error
	SendMessage(ctx context.Context, data *model.SendMessageRequest) error
	StopService(ctx context.Context)
}
