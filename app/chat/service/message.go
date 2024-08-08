package service

import "chat/app/chat/repository"

func NewMessage(messageRepository repository.Message) Message {
	return Message{
		messageRepository: messageRepository,
	}
}

type Message struct {
	messageRepository repository.Message
}
