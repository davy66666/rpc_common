package rabbitmqc

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type ConsumerMetrics struct {
	// 批处理大小分布
	batchSizeHistogram prometheus.Histogram

	// 处理时间分布
	processTimeHistogram prometheus.Histogram

	// 成功/失败计数
	messagesConsumedTotal *prometheus.CounterVec
	messagesFailedTotal   *prometheus.CounterVec

	// 当前正在处理的消息数
	messagesInFlight prometheus.Gauge

	// 延迟时间(消息产生到消费的时间差)
	messageDelayHistogram prometheus.Histogram
}

func NewConsumerMetrics(namespace, subsystem, consumerGroup string) *ConsumerMetrics {
	constLabels := prometheus.Labels{
		"consumer_group": consumerGroup,
	}

	return &ConsumerMetrics{
		batchSizeHistogram: promauto.NewHistogram(
			prometheus.HistogramOpts{
				Namespace:   namespace,
				Subsystem:   subsystem,
				Name:        "batch_size_distribution",
				Help:        "Distribution of batch sizes consumed",
				ConstLabels: constLabels,
				Buckets:     []float64{1, 5, 10, 25, 50, 100, 250, 500, 1000},
			},
		),
		processTimeHistogram: promauto.NewHistogram(
			prometheus.HistogramOpts{
				Namespace:   namespace,
				Subsystem:   subsystem,
				Name:        "process_time_seconds",
				Help:        "Time taken to process a batch of messages",
				ConstLabels: constLabels,
				Buckets:     []float64{0.01, 0.05, 0.1, 0.5, 1, 5, 10, 30},
			},
		),
		messagesConsumedTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace:   namespace,
				Subsystem:   subsystem,
				Name:        "messages_consumed_total",
				Help:        "Total number of messages consumed successfully",
				ConstLabels: constLabels,
			},
			[]string{"topic"},
		),
		messagesFailedTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace:   namespace,
				Subsystem:   subsystem,
				Name:        "messages_failed_total",
				Help:        "Total number of messages failed to process",
				ConstLabels: constLabels,
			},
			[]string{"topic", "error_type"},
		),
		messagesInFlight: promauto.NewGauge(
			prometheus.GaugeOpts{
				Namespace:   namespace,
				Subsystem:   subsystem,
				Name:        "messages_in_flight",
				Help:        "Current number of messages being processed",
				ConstLabels: constLabels,
			},
		),
		messageDelayHistogram: promauto.NewHistogram(
			prometheus.HistogramOpts{
				Namespace:   namespace,
				Subsystem:   subsystem,
				Name:        "message_delay_seconds",
				Help:        "Delay between message production and consumption",
				ConstLabels: constLabels,
				Buckets:     []float64{1, 5, 10, 30, 60, 300, 600, 1800, 3600},
			},
		),
	}
}
