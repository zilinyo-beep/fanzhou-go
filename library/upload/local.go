package upload

import (
	"context"
	"errors"
	"fanzhou-go/api/v1/system"
	"fanzhou-go/internal/app/common/consts"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"strings"
	"time"
)

type Local struct {
}

func (s *Local) Upload(ctx context.Context, file *ghttp.UploadFile) (result system.UploadResponse, err error) {
	if file == nil {
		err = errors.New("文件必须!")
		return
	}
	r := g.RequestFromCtx(ctx)
	host := r.Host
	forwardedHost := r.Header.Get("X-Forwarded-Host")
	if forwardedHost != "" {
		host = forwardedHost
	}
	uri := r.Header.Get("X-Original-URI")
	if uri != "" {
		pos := gstr.PosI(uri, "/api/v1")
		if pos >= 0 {
			uri = gstr.SubStr(uri, 1, pos)
		}
	}
	urlPerfix := fmt.Sprintf("//%s/%s", host, uri)
	p := strings.Trim(consts.UploadPath, "/")
	sp := s.getStaticPath(ctx)
	if sp != "" {
		sp = strings.TrimRight(sp, "/")
	}
	nowData := time.Now().Format("2006-01-02")
	// 包含静态文件夹的路径
	fullDirPath := sp + "/" + p + "/" + nowData
	fileName, err := file.Save(fullDirPath, true)
	if err != nil {
		return
	}
	// 不含静态文件夹的路径
	fullPath := p + "/" + nowData + "/" + fileName

	result = system.UploadResponse{
		Size:     file.Size,
		Path:     fullPath,
		FullPath: urlPerfix + fullPath,
		Name:     file.Filename,
		Type:     file.Header.Get("Content-type"),
	}
	return
}

// 静态文件夹目录
func (s *Local) getStaticPath(ctx context.Context) string {
	value, _ := g.Cfg().Get(ctx, "server.serverRoot")
	if !value.IsEmpty() {
		return value.String()
	}
	return ""
}
