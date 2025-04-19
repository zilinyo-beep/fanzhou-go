package cmd

import (
	"context"
	"fanzhou-go/internal/consts"
	"fanzhou-go/internal/router"
	"fanzhou-go/library/libValidate"
	"fanzhou-go/library/upload"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			//mqueue.InitRabbitMQ()
			//mqueue.StartConsumers()
			g.Log().SetFlags(glog.F_ASYNC | glog.F_TIME_DATE | glog.F_TIME_TIME | glog.F_FILE_LONG)
			g.Log().Info(ctx, gbase64.MustDecodeString(consts.Logo), "Version:", consts.Version)
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				router.R.BindController(ctx, group)
			})
			//重新配置swaggerUI静态页面--start--,若要使用原版gf字段swaggerUI请删除或注释此段
			swaggerPath := g.Cfg().MustGet(ctx, "server.swaggerPath").String()
			if swaggerPath != "" {
				s.BindHookHandler(swaggerPath+"/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
					content := gstr.ReplaceByMap(consts.SwaggerUITemplate, map[string]string{
						`{SwaggerUIDocUrl}`:             g.Cfg().MustGet(ctx, "server.openapiPath").String(),
						`{SwaggerUIDocNamePlaceHolder}`: gstr.TrimRight(fmt.Sprintf(`//%s`, r.Host)),
					})
					r.Response.Write(content)
					r.ExitAll()
				})
			}
			//重新配置swaggerUI静态页面--end--
			enhanceOpenAPIDoc(s)
			//注册相关组件
			register()
			s.Run()
			return nil
		},
	}
)

// 相关组件注册
func register() {
	//注册上传组件
	upload.Register()
	//注册自定义验证规则
	libValidate.Register()
	//执行计划任务
	//task.Run()
}

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Contact: &goai.Contact{
			Name: consts.OpenAPIContactName,
			URL:  consts.OpenAPIContactUrl,
		},
	}
}
