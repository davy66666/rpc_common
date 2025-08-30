package rabbitmqc

import "time"

type (
	QueueOption func(*queueOptions)

	queueOptions struct {
		Size         int
		BatchTimeout time.Duration // 等待批处理超时时间
		RetryTimes   int           // 最大重试次数
	}
)

func WithMaxBatchSize(size int) QueueOption {
	return func(options *queueOptions) {
		options.Size = size
	}
}

func WithMaxBatchTimeout(timeout time.Duration) QueueOption {
	return func(options *queueOptions) {
		options.BatchTimeout = timeout
	}
}

func WithMaxBatchRetryTimes(times int) QueueOption {
	return func(options *queueOptions) {
		options.RetryTimes = times
	}
}
