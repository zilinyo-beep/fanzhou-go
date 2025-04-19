// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskDao is the data access object for the table task.
type TaskDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TaskColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TaskColumns defines and stores column names for the table task.
type TaskColumns struct {
	Id            string // 主键
	TaskId        string // 任务唯一ID（UUID）
	UserId        string // 用户ID
	Params        string // 任务参数（JSON 格式）
	Status        string // 状态：0待处理 1已加速 2已完成 3失败
	QueueType     string // 任务队列类型：0慢队列，1快队列
	TaskType      string // 任务类型：0文生图，1图生图，2图混合，3其他
	RetryCount    string // 重试次数
	AcceleratedAt string // 加速时间（如有）
	CompletedAt   string // 任务完成时间（如有）
	ErrorMsg      string // 任务失败时的错误信息
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// taskColumns holds the columns for the table task.
var taskColumns = TaskColumns{
	Id:            "id",
	TaskId:        "task_id",
	UserId:        "user_id",
	Params:        "params",
	Status:        "status",
	QueueType:     "queue_type",
	TaskType:      "task_type",
	RetryCount:    "retry_count",
	AcceleratedAt: "accelerated_at",
	CompletedAt:   "completed_at",
	ErrorMsg:      "error_msg",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewTaskDao creates and returns a new DAO object for table data access.
func NewTaskDao(handlers ...gdb.ModelHandler) *TaskDao {
	return &TaskDao{
		group:    "default",
		table:    "task",
		columns:  taskColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TaskDao) Columns() TaskColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TaskDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
