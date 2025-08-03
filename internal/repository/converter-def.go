package repository

import (
	repoModel "github.com/Ghaarp/chat-server/internal/repository/chat/model"
	serviceModel "github.com/Ghaarp/chat-server/internal/service/chat/model"
)

type RepoConverter interface {
	ToCreateRequest(req *serviceModel.CreateRequest) *repoModel.CreateRequest
	ToSendMessageRequest(req *serviceModel.SendMessageRequest) *repoModel.SendMessageRequest
}
