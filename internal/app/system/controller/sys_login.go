/*
* @desc:登录
* @company:云南奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2022/4/27 21:52
 */

package controller

import (
	"context"
	"fanzhou-go/api/v1/system"
	commonService "fanzhou-go/internal/app/common/service"
	"fanzhou-go/internal/app/system/model"
	"fanzhou-go/internal/app/system/service"
	"fanzhou-go/library/libUtils"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Login = loginController{}
)

type loginController struct {
	BaseController
}

func (c *loginController) Login(ctx context.Context, req *system.UserLoginReq) (res *system.UserLoginRes, err error) {
	var (
		user        *model.LoginUserRes
		token       string
		permissions []string
		menuList    []*model.UserMenus
	)
	//判断验证码是否正确
	verifyStatus := g.Cfg().MustGet(ctx, "system.verifyStatus").Int()
	if verifyStatus == 1 {
		// 验证码v1版
		if gstr.Trim(req.VerifyCode) == "" {
			err = gerror.New("验证码输入错误")
			return
		}
		if !commonService.Captcha().VerifyString(req.VerifyKey, req.VerifyCode) {
			err = gerror.New("验证码输入错误")
			return
		}
	} else if verifyStatus == 2 {
		// 验证码v2版
		if gstr.Trim(req.VerifyCode) == "" {
			err = gerror.New("人机交互验证失败")
			return
		}
		err = commonService.Captcha().CheckCaptchaV2(ctx, req.VerifyKey, req.VerifyCode, true)
		if err != nil {
			return
		}
	}

	ip := libUtils.GetClientIp(ctx)
	userAgent := libUtils.GetUserAgent(ctx)
	user, err = service.SysUser().GetAdminUserByUsernamePassword(ctx, req)
	if err != nil {
		// 保存登录失败的日志信息
		service.SysLoginLog().Invoke(gctx.New(), &model.LoginLogParams{
			Status:    0,
			Username:  req.Username,
			Ip:        ip,
			UserAgent: userAgent,
			Msg:       err.Error(),
			Module:    "系统后台",
		})
		return
	}
	err = service.SysUser().UpdateLoginInfo(ctx, user.Id, ip)
	if err != nil {
		return
	}
	// 报存登录成功的日志信息
	service.SysLoginLog().Invoke(gctx.New(), &model.LoginLogParams{
		Status:    1,
		Username:  req.Username,
		Ip:        ip,
		UserAgent: userAgent,
		Msg:       "登录成功",
		Module:    "系统后台",
	})
	key := gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword)
	if g.Cfg().MustGet(ctx, "gfToken.multiLogin").Bool() {
		key = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword+ip+userAgent)
	}
	token, err = service.GfToken().GenerateToken(ctx, key, user)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("登录失败，后端服务出现错误")
		return
	}
	//获取用户菜单数据
	menuList, permissions, err = service.SysUser().GetAdminRules(ctx, user.Id)
	if err != nil {
		return
	}
	res = &system.UserLoginRes{
		UserInfo:    user,
		Token:       token,
		MenuList:    menuList,
		Permissions: permissions,
	}
	//用户在线状态保存
	service.SysUserOnline().Invoke(gctx.New(), &model.SysUserOnlineParams{
		UserAgent: userAgent,
		Uuid:      gmd5.MustEncrypt(token),
		Token:     token,
		Username:  user.UserName,
		Ip:        ip,
	})
	return
}

// LoginOut 退出登录
func (c *loginController) LoginOut(ctx context.Context, req *system.UserLoginOutReq) (res *system.UserLoginOutRes, err error) {
	err = service.GfToken().RemoveToken(ctx, service.GfToken().GetRequestToken(g.RequestFromCtx(ctx)))
	return
}
