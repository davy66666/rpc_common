package rabbitmqc

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
)

type RabbitMqConf struct {
	Host          string
	Port          int
	Username      string
	Password      string
	Consumer      string // 消费者组名
	Tag           string // 路由标签
	MaxRetries    int    // 最大重试次数，默认3
	RetryInterval int    `json:"RetryInterval,optional"` // 重试间隔，默认1s
}

type Consumer struct {
	conn           *amqp.Connection
	channel        *amqp.Channel
	conf           RabbitMqConf
	handler        ConsumeHandler
	Topic          string
	ctx            context.Context
	cancel         context.CancelFunc
	wg             sync.WaitGroup
	pool           *threading.RoutineGroup // 协程池
	maxConcur      int                     // 最大并发协程数
	activeRoutines int                     // 当前活跃的协程数
	mu             sync.Mutex              // 用于控制协程数的锁
}

func MustConsumer(c RabbitMqConf, topic string, handler ConsumeHandler) *Consumer {

	if topic == "" {
		panic("invalid Rabbit config")
	}

	url := fmt.Sprintf("amqp://%s:%s@%s:%d", c.Username, c.Password, c.Host, c.Port)
	connection, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Consumer{
		conn: connection,
		//channel:   ch,
		conf:      c,
		handler:   handler,
		Topic:     topic,
		ctx:       ctx,
		cancel:    cancel,
		pool:      threading.NewRoutineGroup(), // 创建协程池
		maxConcur: 10,                          // 最大并发数
	}
}

func (c *Consumer) Start() {

	// 创建通道
	ch, err := c.conn.Channel()
	if err != nil {
		logx.Errorf("创建Channel失败: %v", err)
		panic("failed to create Channel")
	}

	defer ch.Close() // 确保关闭
	//// 设置QoS
	err = ch.Qos(
		1,     // 预取计数
		0,     // 预取大小
		false, // 全局
	)
	if err != nil {
		logx.Errorf("Failed to set QoS: %v", err)
		panic("failed to set QoS")
	}

	for {
		msg, err := ch.Consume(
			c.Topic, // 队列名称
			"",      // 消费者名称
			false,   // 自动确认
			false,   // 排他性
			false,   // 不本地队列
			false,   // 不等待服务器响应
			nil,     // 额外参数
		)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				logx.Info("Consumer context canceled, exiting fetch loop.")
				return
			}
			logx.Errorf("rabbitmq fetch message error: %v", err)
			time.Sleep(time.Second)
			continue
		}

		select {
		case m, ok := <-msg:
			if !ok {
				continue
			}

			ext := MessageExt{
				Delivery:     &m,
				QueueName:    c.Topic,
				ReceivedTime: time.Now(),
			}
			// 启动协程池处理消息
			c.mu.Lock()
			if c.activeRoutines < c.maxConcur {
				c.activeRoutines++ // 增加活跃的协程数
				c.mu.Unlock()

				c.pool.Run(func() {
					defer func() {
						c.mu.Lock()
						c.activeRoutines-- // 协程结束后减少活跃的协程数
						c.mu.Unlock()
					}()

					if err = c.processMessage(&ext); err != nil {
						logx.Errorf("Message process failed after retries: %v", err)
					}

					// 处理成功，确认消息
					ext.Delivery.Ack(false)
				})
			} else {
				c.mu.Unlock()
				// 如果活跃的协程数达到最大并发数，等待所有协程完成
				c.pool.Wait()
			}
		}

	}
}

// processMessage 处理单个消息，包含重试逻辑
func (c *Consumer) processMessage(msg *MessageExt) error {

	var lastErr error
	for i := 0; i < c.conf.MaxRetries; i++ {
		if i > 0 {
			logx.Infof("Retrying message, attempt %d/%d", i+1, c.conf.MaxRetries)
			time.Sleep(time.Duration(c.conf.RetryInterval))
		}

		if err := c.handler.Consume(c.ctx, msg); err == nil {
			return nil
		} else {
			lastErr = err
		}
	}

	return fmt.Errorf("failed after %d retries, last error: %v", c.conf.MaxRetries, lastErr)
}

// Stop 停止消费者
func (c *Consumer) Stop() {
	c.cancel()
	c.pool.Wait() // 等待所有协程完成
	c.wg.Wait()
}

// WaitSignalAndStop 统一优雅退出入口，外部调用
func (c *Consumer) WaitSignalAndStop() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigchan)

	sig := <-sigchan
	logx.Infof("Caught signal %v: shutting down consumer...", sig)
	c.Stop()
}

func initRabbitMQ(c RabbitMqConf) {

	url := fmt.Sprintf("amqp://%s:%s@%s:%d", c.Username, c.Password, c.Host, c.Port)
	conn, err := amqp.DialConfig(url, amqp.Config{
		Heartbeat: 10 * time.Second,
	})
	if err != nil {
		logx.Errorf("rabbitmq restart conn error: %v", err)
	}

	if conn != nil {
		//rabbitConn = conn
		logx.Info("rabbitmq restart conn success")
	}

	// 自动重连
	go func() {
		for {
			reason := <-conn.NotifyClose(make(chan *amqp.Error))
			logx.Errorf("连接关闭: %v, 3秒后重连...", reason)
			time.Sleep(3 * time.Second)
			initRabbitMQ(c)
		}
	}()
}
