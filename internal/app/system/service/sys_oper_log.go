// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"fanzhou-go/api/v1/system"
	"fanzhou-go/internal/app/system/model"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IOperateLog interface {
		OperationLog(r *ghttp.Request)
		Invoke(ctx context.Context, data *model.SysOperLogAdd)
		List(ctx context.Context, req *system.SysOperLogSearchReq) (listRes *system.SysOperLogSearchRes, err error)
		GetByOperId(ctx context.Context, operId uint64) (res *model.SysOperLogInfoRes, err error)
		DeleteByIds(ctx context.Context, ids []uint64) (err error)
		ClearLog(ctx context.Context) (err error)
	}
)

var (
	localOperateLog IOperateLog
)

func OperateLog() IOperateLog {
	if localOperateLog == nil {
		panic("implement not found for interface IOperateLog, forgot register?")
	}
	return localOperateLog
}

func RegisterOperateLog(i IOperateLog) {
	localOperateLog = i
}
