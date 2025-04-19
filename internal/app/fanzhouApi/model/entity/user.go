// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id            int64       `json:"Id"            orm:"id"              description:"主键ID"`
	Phone         string      `json:"Phone"         orm:"phone"           description:"手机号"`
	Nickname      string      `json:"Nickname"      orm:"nickname"        description:"昵称"`
	Coin          int         `json:"Coin"          orm:"coin"            description:"星币余额"`
	DeviceId      string      `json:"DeviceId"      orm:"device_id"       description:"设备唯一ID"`
	DeviceModel   string      `json:"DeviceModel"   orm:"device_model"    description:"设备型号"`
	ClientVersion string      `json:"ClientVersion" orm:"client_version"  description:"客户端版本号"`
	RegisterIp    string      `json:"RegisterIp"    orm:"register_ip"     description:"注册IP"`
	LastLoginIp   string      `json:"LastLoginIp"   orm:"last_login_ip"   description:"最后登录IP"`
	Channel       string      `json:"Channel"       orm:"channel"         description:"注册渠道"`
	Status        int         `json:"Status"        orm:"status"          description:"用户状态 1=正常 0=禁用 -1=封禁"`
	LastLoginTime *gtime.Time `json:"LastLoginTime" orm:"last_login_time" description:"最后登录时间"`
	DeletedAt     *gtime.Time `json:"DeletedAt"     orm:"deleted_at"      description:"删除时间（软删除）"`
	CreatedAt     *gtime.Time `json:"CreatedAt"     orm:"created_at"      description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"UpdatedAt"     orm:"updated_at"      description:"更新时间"`
}
