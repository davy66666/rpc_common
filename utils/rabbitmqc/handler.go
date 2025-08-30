package rabbitmqc

import (
	"context"
	"time"

	"github.com/streadway/amqp"
)

// MessageExt 扩展的消息结构，包含RabbitMQ原始消息和额外元数据
type MessageExt struct {
	Delivery     *amqp.Delivery // 嵌入原始RabbitMQ消息
	QueueName    string         // 来源队列名称
	ConsumerTag  string         // 消费者标签
	ReceivedTime time.Time      // 接收时间
}

type (
	ConsumeHandler interface {
		Consume(ctx context.Context, message *MessageExt) error
	}
	BatchConsumeHandler interface {
		Consume(ctx context.Context, messages ...*MessageExt) error
	}
)
