package router

import (
	"context"
	wechat "fanzhou-go/internal/app/wechat/router"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (router *Router) BindWechatModuleController(ctx context.Context, group *ghttp.RouterGroup) {
	wechat.R.BindController(ctx, group)
}
