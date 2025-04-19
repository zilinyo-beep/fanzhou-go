package rabbitmq

import (
	"github.com/streadway/amqp"
)

// InitQueues 初始化所有队列
func InitQueues(channel *amqp.Channel) error {
	if err := declareFastQueue(channel); err != nil {
		return err
	}
	if err := declareSlowQueue(channel); err != nil {
		return err
	}
	if err := declareDelayQueue(channel); err != nil {
		return err
	}
	return nil
}

// 快队列声明
func declareFastQueue(channel *amqp.Channel) error {
	_, err := channel.QueueDeclare("task_fast_queue", true, false, false, false, nil)
	return err
}

// 慢队列声明（最终消费队列）
func declareSlowQueue(channel *amqp.Channel) error {
	_, err := channel.QueueDeclare("task_slow_queue", true, false, false, false, nil)
	return err
}

// 延迟队列声明（5分钟TTL自动进入慢队列）
func declareDelayQueue(channel *amqp.Channel) error {
	args := amqp.Table{
		"x-dead-letter-exchange":    "",
		"x-dead-letter-routing-key": "task_slow_queue",
		"x-message-ttl":             int32(5 * 60 * 1000),
	}
	_, err := channel.QueueDeclare("task_delay_queue", true, false, false, false, args)
	return err
}
