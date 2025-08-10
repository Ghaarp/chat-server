package converter

import (
	repoModel "github.com/Ghaarp/chat-server/internal/repository/chat/model"
	serviceModel "github.com/Ghaarp/chat-server/internal/service/chat/model"
)

type ChatConverter struct {
}

func CreateConverter() *ChatConverter {
	return &ChatConverter{}
}

func (c *ChatConverter) ToCreateRequest(req *serviceModel.CreateRequest) *repoModel.CreateRequest {
	return &repoModel.CreateRequest{
		Author:   req.Author,
		ChatName: req.ChatName,
		Users:    req.Users,
	}
}

func (c *ChatConverter) ToSendMessageRequest(req *serviceModel.SendMessageRequest) *repoModel.SendMessageRequest {
	return &repoModel.SendMessageRequest{
		From:      req.From,
		Chatid:    req.Chatid,
		Text:      req.Text,
		Timestamp: req.Timestamp,
	}
}
