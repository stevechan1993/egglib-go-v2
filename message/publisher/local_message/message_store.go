package local_message

import (
	"github.com/stevechan1993/egglib-go-v2/core/application"
)

type MessageStore interface {
	AppendMessage(message *application.Message) error
	FindNoPublishedStoredMessages() ([]*application.Message, error)
	FinishPublishStoredMessages(messageIds []int64) error
}
