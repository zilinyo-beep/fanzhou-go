package controller

import (
	"fanzhou-go/internal/app/fanzhouApi/consts"
	service "fanzhou-go/internal/app/fanzhouApi/service"
	"fanzhou-go/utility"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 发送短信验证码
func SendSMS(r *ghttp.Request) {
	var req struct {
		Phone string `p:"phone" v:"required#手机号不能为空"`
	}
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": err.Error()})
		return
	}

	signName := g.Cfg().MustGet(r.GetCtx(), "aliyun.sms.signName").String()

	code := utility.GenerateCode()

	sms := service.GetAliyunSmsService()
	err := sms.SendSMS(req.Phone, signName, map[string]string{"code": code})

	if err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": err.Error()})
		return
	}
	err = g.Redis().SetEX(r.GetCtx(), "sms_code:"+req.Phone, code, consts.SmsRedisTtl)
	if err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": err.Error()})
		return
	}
	r.Response.WriteJsonExit(g.Map{"code": 0, "msg": "发送成功"})
}
func SMSLogin(r *ghttp.Request) {
	var req struct {
		Phone string `p:"phone" v:"required#手机号不能为空"`
		Code  string `p:"Code" v:"required#验证码不能为空"`
	}

	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": err.Error()})
		return
	}
	ctx := r.GetCtx()
	code, _ := g.Redis().Get(ctx, "sms_code:"+req.Phone)

	if code.IsEmpty() || code.String() != req.Code {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "验证码错误或已过期"})
		return
	}

	user, err := service.EnsureUserByPhone(r, req.Phone)
	if err != nil {
		g.Log().Error(ctx, "oneclick-login-error", err)
		r.Response.WriteJsonExit(consts.ServerError)
		return
	}
	token, err := service.GenerateJWT(req.Phone, user.Id)
	if err != nil {
		g.Log().Error(ctx, "jwt-error", err.Error())
		r.Response.WriteJsonExit(consts.ServerError)
		return
	}
	g.Redis().Del(ctx, "sms_code:"+req.Phone)
	r.Response.WriteJsonExit(g.Map{
		"code": 0,
		"msg":  "登录成功",
		"data": g.Map{
			"phone": req.Phone,
			"token": token,
			"coin":  user.Coin,
		},
	})
}

// 客户端拿 verifyToken 调用服务端换手机号
func GetPhoneByToken(r *ghttp.Request) {
	var req struct {
		Token string `p:"token" v:"required#token不能为空"`
	}
	ctx := r.GetCtx()
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": err.Error()})
		return
	}
	authService, err := service.GetAliyunAuthService()
	if err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": err.Error()})
		return
	}

	// 调用方法
	phone, err := authService.GetPhoneByToken("用户提供的token")
	if err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": err.Error()})
		return
	}

	user, err := service.EnsureUserByPhone(r, phone)
	if err != nil {
		g.Log().Error(ctx, "oneclick-login-error", err)
		r.Response.WriteJsonExit(consts.ServerError)
		return
	}
	token, err := service.GenerateJWT(phone, user.Id)
	if err != nil {
		g.Log().Error(ctx, "jwt-error", err.Error())
		r.Response.WriteJsonExit(consts.ServerError)
		return
	}
	r.Response.WriteJsonExit(g.Map{
		"code": 0,
		"msg":  "登录成功",
		"data": g.Map{
			"phone": phone,
			"token": token,
			"coin":  user.Coin,
		},
	})
}
