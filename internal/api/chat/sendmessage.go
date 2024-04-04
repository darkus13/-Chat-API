package chat

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/darkus13/-Chat_API/internal/converter"
	"github.com/darkus13/-Chat_API/internal/repository/chat/model"
	decs "github.com/darkus13/-Chat_API/pkg/chat_v1"
)

func (i *Implementation) SendMessage(ctx context.Context, req *decs.MessageInfo) (*emptypb.Empty, error) {
	err := i.chatService.SendMessage(ctx, (*model.Info)(converter.ToServiceMessage(req)))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
