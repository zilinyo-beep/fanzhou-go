package controller

import (
	"context"
	"fanzhou-go/api/v1/system"
	"fanzhou-go/internal/app/system/service"
)

type bigFile struct{}

var BigFile = new(bigFile)

func (c *bigFile) List(ctx context.Context, req *system.BigFileSearchReq) (res *system.BigFileSearchRes, err error) {
	return service.BigFile().List(ctx, req)
}

// Add 添加系统参数
func (c *bigFile) Add(ctx context.Context, req *system.BigFileAddReq) (res *system.BigFileAddRes, err error) {
	err = service.BigFile().Add(ctx, req, service.Context().GetUserId(ctx))
	return
}

// Get 获取文件信息
func (c *bigFile) Get(ctx context.Context, req *system.BigFileGetReq) (res *system.BigFileGetRes, err error) {
	return service.BigFile().Get(ctx, req.Id)
}

// Edit 修改系统参数
func (c *bigFile) Edit(ctx context.Context, req *system.BigFileEditReq) (res *system.BigFileEditRes, err error) {
	err = service.BigFile().Edit(ctx, req, service.Context().GetUserId(ctx))
	return
}

// Delete 删除系统参数
func (c *bigFile) Delete(ctx context.Context, req *system.BigFileDeleteReq) (res *system.BigFileDeleteRes, err error) {
	err = service.BigFile().Delete(ctx, req.Ids)
	return
}
