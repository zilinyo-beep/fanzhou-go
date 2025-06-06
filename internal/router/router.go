/*
* @desc:路由绑定
* @company:云南奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2022/2/18 16:23
 */

package router

import (
	"context"
	commonRouter "fanzhou-go/internal/app/common/router"
	commonService "fanzhou-go/internal/app/common/service"
	fanzhouApiRouter "fanzhou-go/internal/app/fanzhouApi/router"
	mqueueRouter "fanzhou-go/internal/app/mqueue/router"
	systemRouter "fanzhou-go/internal/app/system/router"
	"fanzhou-go/library/libRouter"
	"github.com/gogf/gf/v2/net/ghttp"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/api", func(group *ghttp.RouterGroup) {
		//跨域处理，安全起见正式环境请注释该行
		group.Middleware(commonService.Middleware().MiddlewareCORS)
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		// 绑定后台路由
		systemRouter.R.BindController(ctx, group)
		// 绑定公共路由
		commonRouter.R.BindController(ctx, group)
		// 测试消息队列用，可以删除
		mqueueRouter.R.BindController(ctx, group)
		fanzhouApiRouter.R.BindController(ctx, group)

		//自动绑定定义的模块
		if err := libRouter.RouterAutoBind(ctx, router, group); err != nil {
			panic(err)
		}
	})
}
