package amqp

import (
	"github.com/hexiaoyun128/gin-base-framework/common"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

var (
	SendConnection *amqp.Connection

	SendChannel *amqp.Channel
	SendQueue   amqp.Queue
)

func AmqpSendDefer() {
	SendChannel.Close()
	SendConnection.Close()
}

//InitAmqpSend
func InitAmqpSend() {
	if common.SendMessageQueueInfo.Type != "amqp" {
		return
	}
	var (
		err error
	)
	SendConnection, err = amqp.Dial(common.SendMessageQueueInfo.Amqp.Url)
	if err != nil {
		common.Logger.Error("amqp Send connect failed", zap.Error(err))
	}
	SendChannel, err = SendConnection.Channel()
	if err != nil {
		common.Logger.Error("amqp Send channel failed", zap.Error(err))
	}
	SendQueue, err = SendChannel.QueueDeclare(
		"block_chain", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		common.Logger.Error("amqp Send queue declare failed", zap.Error(err))
	}
}

func AmqpSender(publishing amqp.Publishing) error {
	err := SendChannel.Publish(
		"",             // exchange
		SendQueue.Name, // routing key
		false,          // mandatory
		false,          // immediate
		publishing)
	return err
}
