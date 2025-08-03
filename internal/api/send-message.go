package chat

import (
	"context"

	generated "github.com/Ghaarp/chat-server/pkg/chat_v1"
)

func (chat *ChatImplementation) SendMessage(ctx context.Context, req *generated.SendMessageRequest) (*generated.SendMessageResponse, error) {

	err := chat.chatService.SendMessage(ctx, chat.serviceConverter.ToSendMessageRequest(req))
	if err != nil {
		return nil, err
	}

	return &generated.SendMessageResponse{}, nil
}
