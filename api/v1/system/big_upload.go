package system

import (
	"fanzhou-go/library/upload_chunk"
	"github.com/gogf/gf/v2/frame/g"
)

// 上传文件
type BigUploadReq struct {
	g.Meta `path:"/bigUpload/upload" tags:"系统后台/后台大文件上传" method:"post" summary:"上传文件"`
	upload_chunk.UploadReq
}

type BigUploadRes struct {
	upload_chunk.UpLoadRes
}

// 上传文件检查
type BigUploadCheckReq struct {
	g.Meta `path:"/bigUpload/upload" tags:"系统后台/后台大文件上传" method:"get" summary:"上传文件检查"`
	upload_chunk.UploadReq
}

type BigUploadCheckRes struct {
	upload_chunk.CheckRes
	Identifier  string `json:"identifier"`  // 标识
	TotalChunks int    `json:"totalChunks"` // 分片总数
}

// 上传文件合并
type BigUploadMergeReq struct {
	g.Meta `path:"/bigUpload/uploadMerge" tags:"系统后台/后台大文件上传" method:"post" summary:"上传文件合并"`
	upload_chunk.UploadReq
}

type BigUploadMergeRes struct {
	upload_chunk.MergeRes
}
