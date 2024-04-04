package chat

import (
	"github.com/darkus13/-Chat_API/internal/repository"
	"github.com/darkus13/-Chat_API/internal/service"
)

type serv struct {
	chatRepository repository.ChatRepository
}

func NewService(chatRepository repository.ChatRepository) service.ChatService {
	return &serv{
		chatRepository: chatRepository,
	}
}
