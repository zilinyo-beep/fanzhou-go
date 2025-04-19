// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Task is the golang structure of table task for DAO operations like Where/Data.
type Task struct {
	g.Meta        `orm:"table:task, do:true"`
	Id            interface{} // 主键
	TaskId        interface{} // 任务唯一ID（UUID）
	UserId        interface{} // 用户ID
	Params        interface{} // 任务参数（JSON 格式）
	Status        interface{} // 状态：0待处理 1已加速 2已完成 3失败
	QueueType     interface{} // 任务队列类型：0慢队列，1快队列
	TaskType      interface{} // 任务类型：0文生图，1图生图，2图混合，3其他
	RetryCount    interface{} // 重试次数
	AcceleratedAt *gtime.Time // 加速时间（如有）
	CompletedAt   *gtime.Time // 任务完成时间（如有）
	ErrorMsg      interface{} // 任务失败时的错误信息
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
