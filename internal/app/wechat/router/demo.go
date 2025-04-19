/*
* @desc:xxxx功能描述
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2023/11/3 16:23
 */

package router

import (
	"context"
	"fanzhou-go/internal/app/wechat/controller"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (router *Router) BindDemoController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/demo", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.Demo,
		)
	})
}
