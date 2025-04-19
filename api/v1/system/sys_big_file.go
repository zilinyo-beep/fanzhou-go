package system

import (
	commonApi "fanzhou-go/api/v1/common"
	"fanzhou-go/internal/app/system/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type BigFileSearchReq struct {
	g.Meta `path:"/bigFile/list" tags:"大文件管理" method:"get" summary:"大文件管理"`
	Name   string `p:"name"` //文件名称
	commonApi.PageReq
}

type BigFileSearchRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.BigFile `json:"list"`
	commonApi.ListRes
}

type BigFileForm struct {
	Name     string `p:"name"  v:"required#文件名称必须"`
	Size     int    `p:"size" v:"required#文件大小必须"`
	Path     string `p:"path" v:"required#文件路径必须"`
	FullPath string `p:"fullPath"`
	MimeType string `p:"mimeType"`
	Source   int    `p:"source"`
	Describe string `p:"describe"`
	Md5      string `p:"md5"`
}

type BigFileGetReq struct {
	g.Meta `path:"/bigFile/get" tags:"系统后台/大文件管理" method:"get" summary:"获取大文件信息"`
	Id     uint64 `p:"id"`
}

type BigFileGetRes struct {
	g.Meta `mime:"application/json"`
	entity.BigFile
}

type BigFileAddReq struct {
	g.Meta `path:"/bigFile/add" tags:"系统后台/大文件管理" method:"post" summary:"大文件新增"`
	BigFileForm
}

type BigFileAddRes struct {
}

type BigFileEditReq struct {
	g.Meta   `path:"/bigFile/edit" tags:"系统后台/大文件管理" method:"put" summary:"大文件新增"`
	Id       uint64 `p:"id" v:"required|min:1#主键ID不能为空|主键ID参数错误"`
	Name     string `p:"name"  v:"required#标题必须"`
	Describe string `p:"describe"`
}

type BigFileEditRes struct {
}

type BigFileDeleteReq struct {
	g.Meta `path:"/bigFile/delete" tags:"系统后台/大文件管理" method:"delete" summary:"大文件删除"`
	Ids    []uint64 `p:"ids"`
}

type BigFileDeleteRes struct {
}
