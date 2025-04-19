/*
* @desc:扶苗工作路由
* @company:云南奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2022/2/18 17:34
 */

package router

import (
	"context"
	"fanzhou-go/internal/app/fanzhouApi/controller"
	"github.com/gogf/gf/v2/net/ghttp"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/auth", func(group *ghttp.RouterGroup) {
		//登录验证拦截
		//service.GfToken().Middleware(group)
		//context拦截器
		//group.Middleware(service.Middleware().Ctx)
		group.POST("/send-code", controller.SendSMS)           // 发短信
		group.POST("/quick-login", controller.GetPhoneByToken) // 验证 token 拿手机号
		group.POST("/code-login", controller.SMSLogin)         // 验证 token 拿手机号

	})
	group.Group("/api/task")
	{
		group.POST("/create", controller.Task.Create)
		group.POST("/accelerate", controller.Task.Accelerate)
	}

}
