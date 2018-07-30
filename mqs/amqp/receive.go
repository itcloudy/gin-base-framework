package amqp

import (
	"fmt"
	"github.com/hexiaoyun128/gin-base-framework/common"
	log "github.com/cihub/seelog"
	"github.com/streadway/amqp"
	"time"
)

var (
	ReceiveConnection *amqp.Connection
	ReceiveChannel    *amqp.Channel
	ReceiveQueue      amqp.Queue
)

func AmqpReceiveDefer() {
	ReceiveChannel.Close()
	ReceiveConnection.Close()
}

//InitAmqpReceive
func InitAmqpReceive() {
	if common.ReceiveMessageQueueInfo.Type != "amqp" {
		return
	}
	var (
		err error
	)
	ReceiveConnection, err = amqp.Dial(common.ReceiveMessageQueueInfo.Amqp.Url)
	if err != nil {
		log.Errorf("amqp Receive connect failed: %s", err)
	}
	ReceiveChannel, err = ReceiveConnection.Channel()
	if err != nil {
		log.Errorf("amqp Receive channel failed: %s", err)
	}
	ReceiveQueue, err = ReceiveChannel.QueueDeclare(
		"block_chain", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)

	if err != nil {
		log.Errorf("amqp Receive queue declare failed: %s", err)
	}
	go amqpReceiver()
}

func amqpReceiver() {
	msgs, err := ReceiveChannel.Consume(
		ReceiveQueue.Name, // queue
		"",                // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	if err != nil {
		log.Errorf("Failed to register a consumer: %s", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			time.Sleep(10000)
			fmt.Println(string(d.Body))

		}
	}()
	<-forever
}
