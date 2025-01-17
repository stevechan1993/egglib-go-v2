package beego

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/stevechan1993/egglib-go-v2/core/application"
	"github.com/stevechan1993/egglib-go-v2/message/receiver/local_message/beego/models"
	"github.com/stevechan1993/egglib-go-v2/transaction/beego"
)

type ReceivedMessageStore struct {
	TransactionContext *beego.TransactionContext
}

func (receivedMessageStore *ReceivedMessageStore) getOrmer() orm.Ormer {
	var ormer orm.Ormer
	if receivedMessageStore.TransactionContext != nil {
		ormer = receivedMessageStore.TransactionContext.Ormer
	} else {
		ormer = orm.NewOrm()
	}
	return ormer
}

func (receivedMessageStore *ReceivedMessageStore) SaveMessage(message *application.Message) error {
	ormer := receivedMessageStore.getOrmer()
	receivedMessageModel := &models.ReceivedMessage{
		Id:          message.MessageId,
		MessageType: message.MessageType,
		MessageBody: message.MessageBody,
		OccurredOn:  message.OccurredOn,
		ReceiveTime: time.Now(),
	}
	if _, err := ormer.Insert(receivedMessageModel); err != nil {
		return err
	}
	return nil
}

func (receivedMessageStore *ReceivedMessageStore) FindMessage(messageId int64) (*application.Message, error) {
	ormer := receivedMessageStore.getOrmer()
	querySeter := ormer.QueryTable("sys_received_messages")
	receivedMessageModel := new(models.ReceivedMessage)
	querySeter.Filter("Id", messageId).One(receivedMessageModel)
	message := &application.Message{
		MessageId:   receivedMessageModel.Id,
		MessageType: receivedMessageModel.MessageType,
		MessageBody: receivedMessageModel.MessageBody,
		OccurredOn:  receivedMessageModel.OccurredOn,
	}
	return message, nil
}
