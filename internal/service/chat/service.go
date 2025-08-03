package chat

import (
	"context"

	repositoryDef "github.com/Ghaarp/chat-server/internal/repository"
	serviceDef "github.com/Ghaarp/chat-server/internal/service"
	"github.com/Ghaarp/chat-server/internal/service/chat/model"
)

type ChatService struct {
	repository repositoryDef.ChatRepository
	converter  repositoryDef.RepoConverter
}

func CreateService(r repositoryDef.ChatRepository, c repositoryDef.RepoConverter) serviceDef.ChatService {
	return &ChatService{
		repository: r,
		converter:  c,
	}
}

func (service *ChatService) Create(ctx context.Context, data *model.CreateRequest) (int64, error) {
	userData := service.converter.ToCreateRequest(data)
	return service.repository.Create(ctx, userData)
}

func (service *ChatService) Delete(ctx context.Context, id int64) error {
	return service.repository.Delete(ctx, id)
}

func (service *ChatService) SendMessage(ctx context.Context, data *model.SendMessageRequest) error {
	return service.repository.SendMessage(ctx, service.converter.ToSendMessageRequest(data))
}

func (service *ChatService) StopService(ctx context.Context) {
	service.repository.ClosePool(ctx)
}
