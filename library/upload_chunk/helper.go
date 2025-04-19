package upload_chunk

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
)

// 解析文件目录，文件前缀及文件后缀
func parseFilePath(filename string) (dir, prefix, suffix string) {
	filenameall := path.Base(filename)                        //不含目录的文件名
	filesuffix := path.Ext(filename)                          // 后缀
	fileprefix := strings.TrimSuffix(filenameall, filesuffix) // 前缀

	dir = strings.TrimRight(filename, filenameall) // 文件目录
	prefix = fileprefix                            // 文件前缀
	suffix = filesuffix                            // 文件后缀
	return
}

// 创建文件
func createFile(fullFilePath string) (ok bool, err error) {
	ok = fileExists(fullFilePath)
	if ok {
		return
	}
	dir, _, _ := parseFilePath(fullFilePath)
	if dir != "" {
		err = mkDir(dir)
		if err != nil {
			return
		}
	}
	newFile, err := os.Create(fullFilePath)
	defer newFile.Close()
	if err != nil {
		return
	}
	return
}

// 创建文件夹
func mkDir(path string) (err error) {
	if !fileExists(path) {
		if err = os.MkdirAll(path, os.ModePerm); err != nil {
			return errors.New(fmt.Sprintf(`os.MkdirAll failed for path "%s" with perm "%d" , err : %v`, path, os.ModePerm, err))
		}
		return nil
	}

	return
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 是否文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// 判断文件或文件夹是否存在
func fileExists(path string) bool {
	if stat, err := os.Stat(path); stat != nil && !os.IsNotExist(err) {
		return true
	}
	return false
}

// 生成分片文件名
func getChunkFilename(identifier, chunkNumber string) string {
	return fmt.Sprintf("uploader-%s.%s", identifier, chunkNumber)
}
