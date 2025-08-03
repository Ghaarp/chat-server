package converter

import (
	serviceDef "github.com/Ghaarp/chat-server/internal/service/chat/model"
	generated "github.com/Ghaarp/chat-server/pkg/chat_v1"
)

type ChatConverter struct {
}

func CreateConverter() *ChatConverter {
	return &ChatConverter{}
}

func (c *ChatConverter) ToCreateRequest(req *generated.CreateRequest) *serviceDef.CreateRequest {
	return &serviceDef.CreateRequest{
		Author:   req.Author,
		ChatName: req.ChatName,
		Users:    req.Users,
	}
}

func (c *ChatConverter) ToSendMessageRequest(req *generated.SendMessageRequest) *serviceDef.SendMessageRequest {
	return &serviceDef.SendMessageRequest{
		From:      req.From,
		ChatId:    req.Chatid,
		Text:      req.Text,
		Timestamp: req.Timestamp.AsTime(),
	}
}
