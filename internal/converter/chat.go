package converter

import (
	modelService "github.com/darkus13/-Chat_API/internal/model"
	decs "github.com/darkus13/-Chat_API/pkg/chat_v1"
)

func ToServiceChat(info *decs.ChatInfo) *modelService.Info {
	return &modelService.Info{
		Owner: info.Owner,
		Users: info.Username,
	}
}

func DeleteFromServices(info *decs.DeleteRequest) *modelService.Info {
	return &modelService.Info{
		ChatID: info.Id,
	}
}

func ToServiceMessage(info *decs.MessageInfo) *modelService.Info {
	return &modelService.Info{
		From:    info.From,
		ChatID:  info.ChatId,
		Content: info.Text,
	}
}
