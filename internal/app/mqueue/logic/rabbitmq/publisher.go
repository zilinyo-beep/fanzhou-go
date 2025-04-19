package rabbitmq

import (
	"context"
	"encoding/json"
	"fanzhou-go/internal/app/fanzhouApi/model/entity"
	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishTask(channel *amqp.Channel, task entity.Task, queueName string) error {
	body, err := json.Marshal(task)
	if err != nil {
		return err
	}

	_, err = channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	return channel.PublishWithContext(context.TODO(), "", queueName, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
}

func PushSlowTask(channel *amqp.Channel, task entity.Task) error {
	return PublishTask(channel, task, "task_slow_queue")
}

func PushFastTask(channel *amqp.Channel, task entity.Task) error {
	return PublishTask(channel, task, "task_fast_queue")
}
