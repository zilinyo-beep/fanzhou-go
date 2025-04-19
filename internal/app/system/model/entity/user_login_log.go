// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLoginLog is the golang structure for table user_login_log.
type UserLoginLog struct {
	Id         int64       `json:"Id"         orm:"id"          description:""`
	Phone      string      `json:"Phone"      orm:"phone"       description:""`
	DeviceId   string      `json:"DeviceId"   orm:"device_id"   description:""`
	DeviceType string      `json:"DeviceType" orm:"device_type" description:""`
	AppVersion string      `json:"AppVersion" orm:"app_version" description:""`
	LoginType  string      `json:"LoginType"  orm:"login_type"  description:""`
	Ip         string      `json:"Ip"         orm:"ip"          description:""`
	CreatedAt  *gtime.Time `json:"CreatedAt"  orm:"created_at"  description:""`
}
