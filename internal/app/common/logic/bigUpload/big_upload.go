/*
* @desc:大文件上传
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/27 16:25
 */

package bigUpload

import (
	"context"
	"fanzhou-go/api/v1/system"
	"fanzhou-go/internal/app/common/service"
	"fanzhou-go/library/upload_chunk"
)

func init() {
	service.RegisterBigUpload(New())
}

func New() service.IBigUpload {
	return &sBigUpload{}
}

type sBigUpload struct{}

// Upload 上传分片文件
func (s *sBigUpload) Upload(ctx context.Context, req *system.BigUploadReq) (res *system.BigUploadRes, err error) {
	uploadChunk := &upload_chunk.UploadChunk{}
	result, err := uploadChunk.Upload(req.UploadReq)
	if err != nil {
		return
	}
	res = new(system.BigUploadRes)
	res.UpLoadRes = *result
	return
}

// UploadCheck 上传文件检查
func (s *sBigUpload) UploadCheck(ctx context.Context, req *system.BigUploadCheckReq) (res *system.BigUploadCheckRes, err error) {
	uploadChunk := &upload_chunk.UploadChunk{}
	result, err := uploadChunk.CheckChunk(req.UploadReq)
	if err != nil {
		return
	}
	//res = new(common.BigUploadCheckRes)
	//res.CheckRes = *result

	res = &system.BigUploadCheckRes{
		CheckRes:    *result,
		Identifier:  req.Identifier,
		TotalChunks: req.TotalChunks,
	}
	return
}

// UploadMerge 上传文件合并
func (s *sBigUpload) UploadMerge(ctx context.Context, req *system.BigUploadMergeReq) (res *system.BigUploadRes, err error) {
	uploadChunk := &upload_chunk.UploadChunk{}
	result, err := uploadChunk.MergeChunk(req.UploadReq)
	if err != nil {
		return
	}
	res = &system.BigUploadRes{
		UpLoadRes: upload_chunk.UpLoadRes{
			BaseRes:     result.BaseRes,
			NeedMerge:   false,
			Identifier:  req.Identifier,
			TotalChunks: req.TotalChunks,
		},
	}

	return
}
