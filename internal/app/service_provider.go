package app

import (
	"context"
	"log"

	"github.com/darkus13/-Chat_API/config"
	"github.com/darkus13/-Chat_API/internal/api/chat"
	"github.com/darkus13/-Chat_API/internal/client/db"
	"github.com/darkus13/-Chat_API/internal/client/db/pg"
	"github.com/darkus13/-Chat_API/internal/closer"
	"github.com/darkus13/-Chat_API/internal/repository"
	chatRepository "github.com/darkus13/-Chat_API/internal/repository/chat"
	"github.com/darkus13/-Chat_API/internal/service"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcCongig config.GRPCConfig

	dbClient       db.Client
	chatService    service.ChatService
	chatRepository repository.ChatRepository

	chatImplm *chat.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcCongig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get gRPC config: %s", err.Error())
		}

		s.grpcCongig = cfg
	}

	return s.grpcCongig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.New
	}
}
