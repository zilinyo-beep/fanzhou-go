// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BigFile is the golang structure for table big_file.
type BigFile struct {
	Id        uint64      `json:"id"        description:""`
	Name      string      `json:"name"      description:"文件名称"`
	Size      int         `json:"size"      description:"文件大小"`
	Path      string      `json:"path"      description:"文件相对路径"`
	FullPath  string      `json:"fullPath"  description:"文件绝对路径"`
	MimeType  string      `json:"mimeType"  description:"文件类型"`
	Source    int         `json:"source"    description:"文件来源  0 - 本地，1 - 腾讯云 2 - 七牛云"`
	Describe  string      `json:"describe"  description:"描述"`
	Md5       string      `json:"md5"       description:"md5"`
	CreatedBy uint64      `json:"createdBy" description:""`
	UpdatedBy uint64      `json:"updatedBy" description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" description:""`
}
