package controller

import (
	"fanzhou-go/internal/app/fanzhouApi/model/entity"
	"fanzhou-go/internal/app/fanzhouApi/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
)

type TaskController struct{}

var Task = TaskController{}

// 提交任务接口：POST /api/task/create
func (c *TaskController) Create(r *ghttp.Request) {
	var input struct {
		UserId int64       `json:"user_id"`
		Params interface{} `json:"params"`
	}
	if err := r.Parse(&input); err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "message": err.Error()})
	}

	task := &entity.Task{
		TaskId:    uuid.New().String(),
		UserId:    uint64(input.UserId),
		Params:    g.NewVar(input.Params).String(),
		QueueType: 0,
		Status:    0,
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
	}

	_, err := service.Task.Create(r.Context(), task)
	if err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "message": "任务创建失败"})
	}

	//go rabbitmq.PublishTask(*task, "task_slow_queue")
	r.Response.WriteJsonExit(g.Map{"code": 0, "message": "任务已提交", "task_id": task.TaskId})
}

// 加速任务接口：POST /api/task/accelerate
func (c *TaskController) Accelerate(r *ghttp.Request) {
	var input struct {
		TaskId string `json:"task_id"`
	}
	if err := r.Parse(&input); err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "message": err.Error()})
	}

	task, err := service.Task.GetByTaskId(r.Context(), input.TaskId)
	if err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "message": "任务不存在"})
	}

	if task.Status != 0 || task.QueueType == 1 {
		r.Response.WriteJsonExit(g.Map{"code": 0, "message": "任务已加速或已处理"})
	}

	task.QueueType = 1
	task.Status = 1
	task.AcceleratedAt = gtime.Now()

	_ = service.Task.UpdateStatus(r.Context(), task.TaskId, 1, "")
	//go rabbitmq.PublishTask(*task, "task_slow_queue")

	r.Response.WriteJsonExit(g.Map{"code": 0, "message": "任务已加速"})
}
