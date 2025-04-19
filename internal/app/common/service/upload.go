// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"fanzhou-go/api/v1/system"

	"github.com/gogf/gf/v2/net/ghttp"
)

type IUpload interface {
	UploadFiles(ctx context.Context, files []*ghttp.UploadFile, checkFileType string, source int) (result system.UploadMultipleRes, err error)
	UploadFile(ctx context.Context, file *ghttp.UploadFile, checkFileType string, source int) (result system.UploadResponse, err error)
	CheckSize(ctx context.Context, checkFileType string, file *ghttp.UploadFile) (err error)
	CheckType(ctx context.Context, checkFileType string, file *ghttp.UploadFile) (err error)
}

var localUpload IUpload

func Upload() IUpload {
	if localUpload == nil {
		panic("implement not found for interface IUpload, forgot register?")
	}
	return localUpload
}

func RegisterUpload(i IUpload) {
	localUpload = i
}
