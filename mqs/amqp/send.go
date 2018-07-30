package amqp

import (
	"github.com/hexiaoyun128/gin-base-framework/common"

	log "github.com/cihub/seelog"
	"github.com/streadway/amqp"
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
		log.Errorf("amqp Send connect failed: %s", err)
	}
	SendChannel, err = SendConnection.Channel()
	if err != nil {
		log.Errorf("amqp Send channel failed: %s", err)
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
		log.Errorf("amqp Send queue declare failed: %s", err)
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
