package local_message

import (
	"github.com/stevechan1993/egglib-go-v2/core/application"
)

type OriginalMessageConverter interface {
	ConvertToMessage(originalMessage interface{}) (*application.Message, error)
}
