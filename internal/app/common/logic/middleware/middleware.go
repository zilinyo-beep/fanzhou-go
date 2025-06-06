/*
* @desc:中间件处理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/28 9:08
 */

package middleware

import (
	"fanzhou-go/internal/app/common/service"
	"github.com/gogf/gf/v2/net/ghttp"
)

func init() {
	service.RegisterMiddleware(New())
}

func New() service.IMiddleware {
	return &sMiddleware{}
}

type sMiddleware struct{}

func (s *sMiddleware) MiddlewareCORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	// you can set options
	//corsOptions.AllowDomain = []string{"goframe.org", "baidu.com"}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}
