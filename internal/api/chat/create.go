package chat

import (
	"context"

	"github.com/darkus13/-Chat_API/internal/converter"
	"github.com/darkus13/-Chat_API/internal/repository/chat/model"
	decs "github.com/darkus13/-Chat_API/pkg/chat_v1"
)

func (i *Implementation) Create(ctx context.Context, req *decs.CreateRequest) (*decs.CreateResponse, error) {

	id, err := i.chatService.Create(ctx, (*model.Info)(converter.ToServiceChat(req.GetInfo())))
	if err != nil {
		return nil, err
	}

	return &decs.CreateResponse{
		Id: id,
	}, nil

}
