package app

import (
	"context"
	"log"

	configDef "github.com/Ghaarp/chat-server/internal/config"
	repositoryDef "github.com/Ghaarp/chat-server/internal/repository"

	pgRepository "github.com/Ghaarp/chat-server/internal/repository/chat"
	pgRepositoryConverter "github.com/Ghaarp/chat-server/internal/repository/chat/converter"

	serviceDef "github.com/Ghaarp/chat-server/internal/service"
	authService "github.com/Ghaarp/chat-server/internal/service/chat"

	chatImplementation "github.com/Ghaarp/chat-server/internal/api"
)

type serviceProvider struct {
	dbConfig     configDef.DBConfig
	serverConfig configDef.ChatConfig

	repository          repositoryDef.ChatRepository
	repositoryConverter repositoryDef.RepoConverter

	service  serviceDef.ChatService
	chatImpl *chatImplementation.ChatImplementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) ChatImplementation(ctx context.Context) *chatImplementation.ChatImplementation {
	if sp.chatImpl == nil {
		sp.chatImpl = chatImplementation.NewChatImplementation(sp.Service(ctx))
	}
	return sp.chatImpl
}

func (sp *serviceProvider) DBConfig() configDef.DBConfig {
	if sp.dbConfig == nil {
		cfg, err := configDef.NewDBConfig()
		if err != nil {
			panic(err)
		}
		sp.dbConfig = cfg
	}
	return sp.dbConfig
}

func (sp *serviceProvider) ServerConfig(configPath string) configDef.ChatConfig {
	if sp.serverConfig == nil {
		err := configDef.Load(configPath)
		if err != nil {
			log.Print("Unable to load .env")
		}

		cfg, err := configDef.NewChatConfig()
		if err != nil {
			panic(err)
		}
		sp.serverConfig = cfg
	}
	return sp.serverConfig
}

func (sp *serviceProvider) RepositoryConverter() repositoryDef.RepoConverter {
	if sp.repositoryConverter == nil {
		sp.repositoryConverter = pgRepositoryConverter.CreateConverter()
	}

	return sp.repositoryConverter
}

func (sp *serviceProvider) Repository(ctx context.Context) repositoryDef.ChatRepository {
	if sp.repository == nil {
		var err error
		sp.repository, err = pgRepository.CreateRepository(ctx, sp.DBConfig().DSN())
		if err != nil {
			log.Fatal(err)
		}
	}
	return sp.repository
}

func (sp *serviceProvider) Service(ctx context.Context) serviceDef.ChatService {
	if sp.service == nil {
		sp.service = authService.CreateService(sp.Repository(ctx), sp.RepositoryConverter())
	}
	return sp.service
}
