package upload_chunk

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"sync"
)

type UploadChunk struct{}

// 检查分片
func (u *UploadChunk) CheckChunk(uploadReq UploadReq) (result *CheckRes, err error) {

	err = uploadReq.MustIdentifier().MustTotalChunks().MustFilename().MustTotalSize().Check()
	if err != nil {
		return
	}
	identifier, totalChunks, filename, totalSize := uploadReq.Identifier, uploadReq.TotalChunks, uploadReq.Filename, uploadReq.TotalSize

	dir, prefix, suffix := parseFilePath(filename)
	_, _, _ = dir, prefix, suffix

	if !strings.Contains(suffix, ".") {
		err = errors.New("文件名解析错误")
		return
	}

	result = &CheckRes{}

	// 秒传
	resultFilePath := u.Tmp() + identifier + suffix
	if fileExists(resultFilePath) {
		result.SkipUpload = true
		result.Url = resultFilePath
		result.Filename = filename
		result.TotalSize = totalSize
		result.RelativePath()
		return
	}

	// 断点续传
	for i := 1; i <= totalChunks; i++ {
		chunkFilePath := u.chunkPath(identifier, strconv.FormatInt(int64(i), 10))
		if fileExists(chunkFilePath) {
			result.Uploaded = append(result.Uploaded, i)
		}
	}

	return
}

// 合并文件
func (u *UploadChunk) MergeChunk(uploadReq UploadReq) (result *MergeRes, err error) {
	err = uploadReq.MustIdentifier().MustTotalChunks().MustTotalSize().MustFilename().Check()
	if err != nil {
		return
	}
	identifier, totalChunks, totalSize, filename := uploadReq.Identifier, uploadReq.TotalChunks, uploadReq.TotalSize, uploadReq.Filename

	_, _, suffix := parseFilePath(filename)
	if !strings.Contains(suffix, ".") {
		return nil, errors.New("文件名解析错误")
	}

	// 合并后的文件
	resultFilePath := u.Tmp() + identifier + suffix
	ok, err := createFile(resultFilePath)
	if err != nil {
		return
	}
	result = new(MergeRes)
	// 文件已存在
	if ok {
		result.Url = resultFilePath
		result.Filename = filename
		result.TotalSize = totalSize
		result.RelativePath()
		return
	}

	// 检查分片文件是否完整
	ok, err = u.checkChunkAll(identifier, totalChunks, totalSize)
	if err != nil {
		err = fmt.Errorf("分片文件检查错误:%s", err)
		return
	}

	if !ok {
		err = errors.New("分片文件不完整")
		return
	}

	var chunkSize int64
	var wg sync.WaitGroup
	ch := make(chan struct{}, 10)
	for i := 1; i <= totalChunks; i++ {
		// 分片文件
		filePath := u.chunkPath(identifier, fmt.Sprintf("%d", i))
		if chunkSize == 0 {
			fi, _ := os.Stat(filePath)
			if chunkSize = fi.Size(); chunkSize == 0 {
				err = errors.New("分片文件大小为0")
				return
			}
		}
		ch <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer func() {
				<-ch
				wg.Done()
			}()
			uploadChunk := &UploadChunk{}
			err = uploadChunk.mergeFile(filePath, resultFilePath, chunkSize*int64(i-1))
			if err != nil {
				log.Println(err)
				return
			}
		}(i)
	}
	wg.Wait()

	result.Url = resultFilePath
	result.Filename = filename
	result.TotalSize = totalSize
	result.RelativePath()
	return
}

// 上传分片文件
func (u *UploadChunk) Upload(uploadReq UploadReq) (result *UpLoadRes, err error) {
	err = uploadReq.MustTotalChunks().MustChunkNumber().MustIdentifier().MustFile().MustTotalSize().MustFilename().Check()
	if err != nil {
		return
	}
	totalChunks, chunkNumber, identifier, upFile, filename := uploadReq.TotalChunks, uploadReq.ChunkNumber, uploadReq.Identifier, uploadReq.File, uploadReq.Filename
	var fullFilePath string
	if totalChunks > 1 {
		// 分片文件路径
		fullFilePath = u.chunkPath(identifier, strconv.Itoa(chunkNumber))
	} else {
		_, _, suffix := parseFilePath(filename)
		if !strings.Contains(suffix, ".") {
			return nil, errors.New("文件名解析错误")
		}
		fullFilePath = u.Tmp() + identifier + suffix
	}

	// 创建文件
	ok, err := createFile(fullFilePath)

	if err != nil {
		return
	}
	if ok {
		// 文件已经存在
		result = &UpLoadRes{
			BaseRes: BaseRes{
				Filename:  uploadReq.Filename,
				TotalSize: uploadReq.TotalSize,
				Url:       fullFilePath,
			},
			NeedMerge:   totalChunks > 1,
			Identifier:  uploadReq.Identifier,
			TotalChunks: uploadReq.TotalChunks,
		}
		result.RelativePath()
		return
	}

	// 打开分片文件
	file, err := os.OpenFile(fullFilePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	// 文件偏移量
	var fi os.FileInfo
	fi, err = os.Stat(fullFilePath)
	if err != nil {
		return
	}
	start := fi.Size()
	// 写入分片文件
	_, err = u.writeFile(upFile, start, file, start)
	if err != nil {
		return
	}
	result = &UpLoadRes{
		BaseRes: BaseRes{
			Filename:  uploadReq.Filename,
			TotalSize: uploadReq.TotalSize,
			Url:       fullFilePath,
		},
		NeedMerge:   totalChunks > 1,
		Identifier:  uploadReq.Identifier,
		TotalChunks: uploadReq.TotalChunks,
	}
	result.RelativePath()
	return
}

func (u *UploadChunk) Tmp() string {
	return g.Cfg().MustGet(gctx.New(), "server.serverRoot").String() + "/big_file/" // 文件保存目录
}

// 合并文件
func (u *UploadChunk) mergeFile(chunkFile, mergeFile string, offset int64) (err error) {

	// 合并后的文件
	file, err := os.OpenFile(mergeFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = file.Seek(offset, 0)
	if err != nil {
		return
	}
	// 分片文件
	chunkFileObj, err := os.Open(chunkFile)
	if err != nil {
		return
	}
	defer chunkFileObj.Close()
	// 写入数据
	data := make([]byte, 1024, 1024)

	for {
		tal, e := chunkFileObj.Read(data)
		if e == io.EOF {
			chunkFileObj.Close()
			os.Remove(chunkFile)
			break
		}
		_, e = file.Write(data[:tal])
		if e != nil {
			return e
		}
	}
	return
}

// 检查分片文件是否完整
func (u *UploadChunk) checkChunkAll(identifier string, totalChunks int, totalSize int64) (ok bool, err error) {
	if identifier == "" || totalChunks == 0 {
		return false, errors.New("checkChunkAll 参数错误")
	}
	var _totalSize int64
	for i := 1; i <= totalChunks; i++ {
		filePath := u.chunkPath(identifier, fmt.Sprintf("%d", i))
		fi, e := os.Stat(filePath)
		if e != nil {
			return false, e
		}
		_totalSize += fi.Size()
	}

	return _totalSize == totalSize, nil
}

// 获取分片文件路径
func (u *UploadChunk) chunkPath(identifier string, chunkNumber string) string {
	return fmt.Sprintf("%s%s", u.Tmp(), getChunkFilename(identifier, chunkNumber))
}

// 检查文件完整性
func (u *UploadChunk) verifyFileSize(fullFilePath string, size int64) (ok bool, err error) {
	fi, err := os.Stat(fullFilePath)
	if err != nil {
		return
	}
	if fi.Size() == size {
		return true, nil
	}
	return false, nil
}

// 写入文件
func (u *UploadChunk) writeFile(upfile multipart.File, upSeek int64, file *os.File, fSeek int64) (result int, err error) {
	// 上传文件大小记录
	fileSzie := 0
	// 设置上传偏移量
	upfile.Seek(upSeek, 0)
	// 设置文件偏移量
	file.Seek(fSeek, 0)
	data := make([]byte, 1024, 1024)
	for {
		total, e := upfile.Read(data)
		if e == io.EOF {
			// 文件复制完毕
			break
		}
		l, e := file.Write(data[:total])
		if e != nil {
			return 0, errors.New("文件上传失败")
		}
		// 记录上传长度
		fileSzie += l
	}
	return fileSzie, nil
}
