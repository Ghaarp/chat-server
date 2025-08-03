package chat

import (
	"github.com/Ghaarp/chat-server/internal/service"
	converter "github.com/Ghaarp/chat-server/internal/service/chat/converter"
	generated "github.com/Ghaarp/chat-server/pkg/chat_v1"
)

type ChatImplementation struct {
	generated.UnimplementedChatV1Server
	chatService      service.ChatService
	serviceConverter service.ServiceConverter
}

func NewChatImplementation(chatService service.ChatService) *ChatImplementation {
	return &ChatImplementation{
		chatService:      chatService,
		serviceConverter: converter.CreateConverter(),
	}
}
