// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Task is the golang structure for table task.
type Task struct {
	Id            uint64      `json:"Id"            orm:"id"             description:"主键"`
	TaskId        string      `json:"TaskId"        orm:"task_id"        description:"任务唯一ID（UUID）"`
	UserId        uint64      `json:"UserId"        orm:"user_id"        description:"用户ID"`
	Params        string      `json:"Params"        orm:"params"         description:"任务参数（JSON 格式）"`
	Status        int         `json:"Status"        orm:"status"         description:"状态：0待处理 1已加速 2已完成 3失败"`
	QueueType     int         `json:"QueueType"     orm:"queue_type"     description:"任务队列类型：0慢队列，1快队列"`
	TaskType      int         `json:"TaskType"      orm:"task_type"      description:"任务类型：0文生图，1图生图，2图混合，3其他"`
	RetryCount    int         `json:"RetryCount"    orm:"retry_count"    description:"重试次数"`
	AcceleratedAt *gtime.Time `json:"AcceleratedAt" orm:"accelerated_at" description:"加速时间（如有）"`
	CompletedAt   *gtime.Time `json:"CompletedAt"   orm:"completed_at"   description:"任务完成时间（如有）"`
	ErrorMsg      string      `json:"ErrorMsg"      orm:"error_msg"      description:"任务失败时的错误信息"`
	CreatedAt     *gtime.Time `json:"CreatedAt"     orm:"created_at"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"UpdatedAt"     orm:"updated_at"     description:"更新时间"`
}
