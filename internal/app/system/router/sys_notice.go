// ==========================================================================
// GFast自动生成router操作代码。
// 生成日期：2023-11-09 11:41:17
// 生成路径: internal/app/system/router/sys_notice.go
// 生成人：gfast
// desc:通知公告
// company:云南奇讯科技有限公司
// ==========================================================================

package router

import (
	"context"

	"fanzhou-go/internal/app/system/controller"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (router *Router) BindSysNoticeController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/sysNotice", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.SysNotice,
		)
	})
}
