package rabbitmqc

import (
	"context"
	"fmt"
	"log"

	//"reflect"
	//"strings"
	"time"

	"github.com/streadway/amqp"
	//"github.com/zeromicro/go-zero/core/logx"
	//xerror "github.com/davy66666/rpc_common/utils/errorx"
)

type BatchConsumer struct {
	conn       *amqp.Connection
	channel    *amqp.Channel
	conf       RabbitMqConf
	handler    BatchConsumeHandler
	topic      string
	ctx        context.Context
	cancelFunc context.CancelFunc
	metrics    *ConsumerMetrics
}

func NewRabbitConnect(conf RabbitMqConf) *amqp.Connection {

	url := fmt.Sprintf("amqp://%s:%s@%s:%d", conf.Username, conf.Password, conf.Host, conf.Port)
	conn, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}

	return conn
}

func MustBatchConsumer(c RabbitMqConf, topic string, handler BatchConsumeHandler, opts ...QueueOption) *BatchConsumer {
	option := queueOptions{
		Size:         500,
		RetryTimes:   3,
		BatchTimeout: 5 * time.Second,
	}
	for _, opt := range opts {
		opt(&option)
	}

	url := fmt.Sprintf("amqp://%s:%s@%s:%d", c.Username, c.Password, c.Host, c.Port)
	connection, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	// 创建监控指标
	metrics := NewConsumerMetrics(
		"rocketmq",
		"consumer",
		"",
	)
	conn := newBatchConsumer(connection, c, topic, handler, option, metrics)
	return &BatchConsumer{
		conn:       connection,
		channel:    conn,
		conf:       c,
		handler:    handler,
		topic:      topic,
		ctx:        ctx,
		cancelFunc: cancelFunc,
		metrics:    metrics,
	}
}

func newBatchConsumer(connection *amqp.Connection, c RabbitMqConf, topic string, handler BatchConsumeHandler, option queueOptions, metrics *ConsumerMetrics) *amqp.Channel {

	// 创建通道
	ch, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	// 设置QoS
	err = ch.Qos(
		1,     // 预取计数
		0,     // 预取大小
		false, // 全局
	)
	if err != nil {
		log.Fatalf("Failed to set QoS: %v", err)
	}

	// 设置QoS，限制未确认消息数量
	err = ch.Qos(
		10,    // 预取计数
		0,     // 预取大小
		false, // 应用于全局
	)
	if err != nil {
		log.Fatalf("Failed to set QoS: %v", err)
	}

	return ch
}

//func (c *BatchConsumer) Start() {
//	if err := c.conn.Start(); err != nil {
//		panic(err)
//	}
//	logx.Infof("RocketMQ batch consumer for topic [%s] started.", c.topic)
//}

//func (c *BatchConsumer) Stop() {
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//
//	done := make(chan struct{})
//	go func() {
//		c.cancelFunc()
//		err := c.conn.Shutdown()
//		if err != nil {
//			logx.Error(err)
//		}
//		close(done)
//	}()
//
//	select {
//	case <-done:
//		return
//	case <-ctx.Done():
//		return
//	}
//}
