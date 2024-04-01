package chat

import (
	"github.com/darkus13/-Chat_API/internal/service"
	decs "github.com/darkus13/-Chat_API/pkg/chat_v1"
)

type Implementation struct {
	decs.UnimplementedChatV1Server
	chatService service.ChatService
}

func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
