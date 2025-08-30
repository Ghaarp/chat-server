package app

import (
	"context"
	"flag"
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
	dbConfig   configDef.DBConfig
	grpcConfig configDef.Config
	httpConfig configDef.Config

	repository          repositoryDef.ChatRepository
	repositoryConverter repositoryDef.RepoConverter

	service  serviceDef.ChatService
	chatImpl *chatImplementation.ChatImplementation
}

func newServiceProvider() *serviceProvider {
	var configPath string
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")

	err := configDef.Load(configPath)
	if err != nil {
		log.Print("Unable to load .env")
	}
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

func (sp *serviceProvider) GRPCConfig() configDef.Config {
	if sp.grpcConfig == nil {
		cfg, err := configDef.NewGrpcConfig()
		if err != nil {
			panic(err)
		}
		sp.grpcConfig = cfg
	}
	return sp.grpcConfig
}

func (sp *serviceProvider) HttpConfig() configDef.Config {
	if sp.httpConfig == nil {
		cfg, err := configDef.NewHttpConfig()
		if err != nil {
			panic(err)
		}
		sp.httpConfig = cfg
	}
	return sp.httpConfig
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
