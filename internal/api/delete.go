package chat

import (
	"context"

	generated "github.com/Ghaarp/chat-server/pkg/chat_v1"
)

func (chat *ChatImplementation) Delete(ctx context.Context, in *generated.DeleteRequest) (*generated.DeleteResponse, error) {

	err := chat.chatService.Delete(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &generated.DeleteResponse{}, nil
}
