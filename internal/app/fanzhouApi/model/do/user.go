// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta        `orm:"table:user, do:true"`
	Id            interface{} // 主键ID
	Phone         interface{} // 手机号
	Nickname      interface{} // 昵称
	Coin          interface{} // 星币余额
	DeviceId      interface{} // 设备唯一ID
	DeviceModel   interface{} // 设备型号
	ClientVersion interface{} // 客户端版本号
	RegisterIp    interface{} // 注册IP
	LastLoginIp   interface{} // 最后登录IP
	Channel       interface{} // 注册渠道
	Status        interface{} // 用户状态 1=正常 0=禁用 -1=封禁
	LastLoginTime *gtime.Time // 最后登录时间
	DeletedAt     *gtime.Time // 删除时间（软删除）
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
