package chat

import (
	"context"

	generated "github.com/Ghaarp/chat-server/pkg/chat_v1"
)

func (chat *ChatImplementation) Create(ctx context.Context, req *generated.CreateRequest) (*generated.CreateResponse, error) {

	id, err := chat.chatService.Create(ctx, chat.serviceConverter.ToCreateRequest(req))
	if err != nil {
		return nil, err
	}

	return &generated.CreateResponse{
		Id: id,
	}, nil
}
