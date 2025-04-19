// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDao is the data access object for the table user.
type UserDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UserColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UserColumns defines and stores column names for the table user.
type UserColumns struct {
	Id            string // 主键ID
	Phone         string // 手机号
	Nickname      string // 昵称
	Coin          string // 星币余额
	DeviceId      string // 设备唯一ID
	DeviceModel   string // 设备型号
	ClientVersion string // 客户端版本号
	RegisterIp    string // 注册IP
	LastLoginIp   string // 最后登录IP
	Channel       string // 注册渠道
	Status        string // 用户状态 1=正常 0=禁用 -1=封禁
	LastLoginTime string // 最后登录时间
	DeletedAt     string // 删除时间（软删除）
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// userColumns holds the columns for the table user.
var userColumns = UserColumns{
	Id:            "id",
	Phone:         "phone",
	Nickname:      "nickname",
	Coin:          "coin",
	DeviceId:      "device_id",
	DeviceModel:   "device_model",
	ClientVersion: "client_version",
	RegisterIp:    "register_ip",
	LastLoginIp:   "last_login_ip",
	Channel:       "channel",
	Status:        "status",
	LastLoginTime: "last_login_time",
	DeletedAt:     "deleted_at",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao(handlers ...gdb.ModelHandler) *UserDao {
	return &UserDao{
		group:    "default",
		table:    "user",
		columns:  userColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserDao) Columns() UserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
