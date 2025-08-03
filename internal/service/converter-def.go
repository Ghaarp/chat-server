package service

import (
	"github.com/Ghaarp/chat-server/internal/service/chat/model"
	generated "github.com/Ghaarp/chat-server/pkg/chat_v1"
)

type ServiceConverter interface {
	ToCreateRequest(req *generated.CreateRequest) *model.CreateRequest
	ToSendMessageRequest(req *generated.SendMessageRequest) *model.SendMessageRequest
}
