package mqueue

import (
	"fanzhou-go/internal/app/fanzhouApi/model/entity"
	"fanzhou-go/internal/app/mqueue/logic/rabbitmq"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func TestRabbitMQPublishAndConsume(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := rabbitmq.InitRabbitMQ()
		t.AssertNil(err)
		testTask := entity.Task{
			TaskId: "test123",
			UserId: 9999,
			Params: `{"name":"test"}`,
		}
		// 发送消息
		err = rabbitmq.PublishTask(testTask, "task_slow_queue")
		t.AssertNil(err)
	})
}
