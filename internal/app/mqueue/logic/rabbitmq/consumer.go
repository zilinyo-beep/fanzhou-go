package rabbitmq

import (
	"context"
	"encoding/json"
	"fanzhou-go/internal/app/fanzhouApi/model/entity"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/streadway/amqp"
)

func StartConsumers(channel *amqp.Channel) {
	startConsumer(channel, "task_fast_queue", handleFastTask)
	startConsumer(channel, "task_slow_queue", handleSlowTask)
}

func startConsumer(channel *amqp.Channel, queueName string, handler func(entity.Task)) {
	_, err := channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	msgs, err := channel.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	go func() {
		for msg := range msgs {
			var task entity.Task
			if err := json.Unmarshal(msg.Body, &task); err == nil {
				fmt.Println("[RECEIVED]", task)
				handler(task)
			} else {
				fmt.Println("[UNMARSHAL ERROR]", err)
			}
		}
	}()
}

func handleFastTask(task entity.Task) {
	fmt.Println("FAST TASK")
	g.Log().Infof(context.Background(), "[FAST] 处理任务: %+v", task)
}

func handleSlowTask(task entity.Task) {
	fmt.Println("SLOW TASK")
	g.Log().Infof(context.Background(), "[SLOW] 处理任务: %+v", task)
}
