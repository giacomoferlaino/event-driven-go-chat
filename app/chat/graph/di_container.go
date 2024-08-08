package graph

import (
	"chat/app/chat/repository"
	"chat/app/chat/service"
)

type diContainer struct {
	MessageService service.Message
}

func newDIContainer() diContainer {
	messageRepository := repository.NewMessage()
	messageService := service.NewMessage(messageRepository)

	return diContainer{
		MessageService: messageService,
	}
}
