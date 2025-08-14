package rabbitmqc

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
)

func MustProducer(c RabbitMqConf) *amqp.Connection {

	url := fmt.Sprintf("amqp://%s:%s@%s:%d", c.Username, c.Password, c.Host, c.Port)
	connection, err := amqp.DialConfig(url, amqp.Config{
		Heartbeat: 10 * time.Second,
	})
	if err != nil {
		logx.Errorf("failed to create rabbitmq producer: %v", err)
		panic(err)
	}

	logx.Infof("rabbitmq producer started for %v", url)
	return connection
}
