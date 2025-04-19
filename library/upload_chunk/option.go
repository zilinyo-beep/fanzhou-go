package upload_chunk

import (
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
)

type checkErr func() error

type UploadReq struct {
	ChunkNumber      int    `json:"chunkNumber"`      // 分片序列号
	ChunkSize        int    `json:"chunkSize"`        // 分片大小
	CurrentChunkSize int    `json:"currentChunkSize"` // 当前分片大小
	TotalSize        int64  `json:"totalSize"`        // 文件总大小
	Identifier       string `json:"identifier"`       // 标识
	Filename         string `json:"filename"`         // 文件名
	RelativePath     string `json:"relativePath"`     //文件夹上传的时候文件的相对路径属性
	TotalChunks      int    `json:"totalChunks"`      // 分片总数
	File             multipart.File
	FileHeader       *multipart.FileHeader

	verification map[interface{}][]checkErr
}

func (u *UploadReq) verificationInit() {
	if u.verification == nil {
		u.verification = make(map[interface{}][]checkErr)
	}
}

func (u *UploadReq) MustChunkNumber() *UploadReq {
	u.verificationInit()
	u.verification["ChunkNumber"] = append(u.verification["ChunkNumber"], func() error {
		if u.ChunkNumber == 0 {
			return errors.New("ChunkNumber不能为空")
		}
		return nil
	})

	return u
}

func (u *UploadReq) MustChunkSize() *UploadReq {
	u.verificationInit()
	u.verification["ChunkSize"] = append(u.verification["ChunkSize"], func() error {
		if u.ChunkSize == 0 {
			return errors.New("ChunkSize不能为空")
		}
		return nil
	})

	return u
}

func (u *UploadReq) MustCurrentChunkSize() *UploadReq {
	u.verificationInit()
	u.verification["CurrentChunkSize"] = append(u.verification["CurrentChunkSize"], func() error {
		if u.CurrentChunkSize == 0 {
			return errors.New("CurrentChunkSize不能为空")
		}
		return nil
	})
	return u
}

func (u *UploadReq) MustTotalSize() *UploadReq {
	u.verificationInit()
	u.verification["TotalSize"] = append(u.verification["TotalSize"], func() error {
		if u.TotalSize == 0 {
			return errors.New("TotalSize不能为空")
		}
		return nil
	})
	return u
}

func (u *UploadReq) MustIdentifier() *UploadReq {
	u.verificationInit()
	u.verification["Identifier"] = append(u.verification["Identifier"], func() error {
		if u.Identifier == "" {
			return errors.New("Identifier不能为空")
		}
		return nil
	})
	return u
}

func (u *UploadReq) MustFilename() *UploadReq {
	u.verificationInit()
	u.verification["Filename"] = append(u.verification["Filename"], func() error {
		if u.Filename == "" {
			return errors.New("Filename不能为空")
		}
		return nil
	})
	return u
}

func (u *UploadReq) MustTotalChunks() *UploadReq {
	u.verificationInit()
	u.verification["TotalChunks"] = append(u.verification["TotalChunks"], func() error {
		if u.TotalChunks == 0 {
			return errors.New("TotalChunks不能为空")
		}
		return nil
	})
	return u
}

func (u *UploadReq) MustFile() *UploadReq {
	u.verificationInit()
	u.verification["File"] = append(u.verification["File"], func() error {
		if u.File == nil || u.FileHeader == nil {
			return errors.New("File不能为空")
		}
		return nil
	})
	return u
}

func (u *UploadReq) Check() (err error) {
	for _, item := range u.verification {
		for _, fn := range item {
			err = fn()
			if err != nil {
				return
			}

		}
	}
	return
}

func (u *UploadReq) Bind(r *http.Request) (err error) {
	if r.Header.Get("Content-Type") == "application/json" {
		l := r.ContentLength
		if l == 0 {
			return
		}
		buf := make([]byte, l)
		_, err = r.Body.Read(buf)

		if err != nil && err != io.EOF {
			return
		}
		err = json.Unmarshal(buf, u)
		if err != nil {
			return
		}

	} else {
		chunkNumber := r.FormValue("chunkNumber")
		if chunkNumber != "" {
			u.ChunkNumber, _ = strconv.Atoi(chunkNumber)
		}

		chunkSize := r.FormValue("chunkSize")
		if chunkSize != "" {
			u.ChunkSize, _ = strconv.Atoi(chunkSize)
		}

		currentChunkSize := r.FormValue("currentChunkSize")
		if currentChunkSize != "" {
			u.CurrentChunkSize, _ = strconv.Atoi(currentChunkSize)
		}

		totalSize := r.FormValue("totalSize")
		if totalSize != "" {
			u.TotalSize, _ = strconv.ParseInt(totalSize, 10, 64)
		}

		identifier := r.FormValue("identifier")
		if identifier != "" {
			u.Identifier = identifier
		}

		filename := r.FormValue("filename")
		if filename != "" {
			u.Filename = filename
		}

		relativePath := r.FormValue("relativePath")
		if relativePath != "" {
			u.RelativePath = relativePath
		}

		totalChunks := r.FormValue("totalChunks")
		if totalChunks != "" {
			u.TotalChunks, _ = strconv.Atoi(totalChunks)
		}
	}

	u.File, u.FileHeader, _ = r.FormFile(FileName)
	return
}

type BaseRes struct {
	Filename  string `json:"filename"`  // 文件名
	TotalSize int64  `json:"totalSize"` // 文件总大小
	Url       string `json:"url"`       // 上传文件路径
}

// 返回前端的相对路径
func (b *BaseRes) RelativePath() {
	_, prefix, suffix := parseFilePath(b.Url)
	b.Url = RelativePath + prefix + suffix
}

type UpLoadRes struct {
	BaseRes
	NeedMerge   bool   `json:"needMerge"`   // 是否合并文件
	Identifier  string `json:"identifier"`  // 标识
	TotalChunks int    `json:"totalChunks"` // 分片总数
}

type CheckRes struct {
	SkipUpload bool  `json:"skipUpload"` // 秒传
	Uploaded   []int `json:"uploaded"`   // 已上传过的分片
	BaseRes
}

func (r *CheckRes) HasUploaded(i int) (ok bool) {
	for _, item := range r.Uploaded {
		if item == i {
			return true
		}
	}
	return
}

func (r *CheckRes) HasFirst() (ok bool) {
	return r.HasUploaded(1)
}

type MergeRes struct {
	BaseRes
}
