package router

import (
	"context"
	websocketRouter "fanzhou-go/internal/app/websocket/router"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (router *Router) BindApiModuleController(ctx context.Context, group *ghttp.RouterGroup) {
	websocketRouter.R.BindController(ctx, group)
}
