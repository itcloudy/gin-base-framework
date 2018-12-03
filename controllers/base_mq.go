package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/itcloudy/gin-base-framework/common"
	amqp2 "github.com/itcloudy/gin-base-framework/mqs/amqp"
	"github.com/streadway/amqp"
)

func MessageQueue(c *gin.Context) {
	m := amqp.Publishing{}

	m.Body = []byte(c.Query("msg"))
	m.ContentType = "text/plain"
	common.GenResponse(c, common.SUCCESSED, amqp2.AmqpSender(m), "message queue")
}
