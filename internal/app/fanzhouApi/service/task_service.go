package service

import (
	"context"
	"errors"
	"fanzhou-go/internal/app/fanzhouApi/dao"
	"fanzhou-go/internal/app/fanzhouApi/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type taskService struct{}

var Task = taskService{}

func (s *taskService) Create(ctx context.Context, task *entity.Task) (int64, error) {
	result, err := dao.Task.Ctx(ctx).Data(task).Insert()
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (s *taskService) GetByTaskId(ctx context.Context, taskId string) (*entity.Task, error) {
	var task *entity.Task
	err := dao.Task.Ctx(ctx).Where("task_id", taskId).Scan(&task)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, errors.New("任务不存在")
	}
	return task, nil
}

func (s *taskService) UpdateStatus(ctx context.Context, taskId string, status int, errorMsg string) error {
	_, err := dao.Task.Ctx(ctx).Where("task_id", taskId).Data(g.Map{
		"status":     status,
		"error_msg":  errorMsg,
		"updated_at": time.Now(),
	}).Update()
	return err
}
