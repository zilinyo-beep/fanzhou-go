/*
* @desc:xxxx功能描述
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/11/3 10:32
 */

package controller

import (
	"context"
	"fanzhou-go/api/v1/system"
	"fanzhou-go/internal/app/system/service"
	"fanzhou-go/library/libUtils"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

var Personal = new(personalController)

type personalController struct {
}

func (c *personalController) GetPersonal(ctx context.Context, req *system.PersonalInfoReq) (res *system.PersonalInfoRes, err error) {
	res, err = service.Personal().GetPersonalInfo(ctx, req)
	return
}

func (c *personalController) EditPersonal(ctx context.Context, req *system.PersonalEditReq) (res *system.PersonalEditRes, err error) {
	ip := libUtils.GetClientIp(ctx)
	userAgent := libUtils.GetUserAgent(ctx)
	res = new(system.PersonalEditRes)
	res.UserInfo, err = service.Personal().EditPersonal(ctx, req)
	if err != nil {
		return
	}
	key := gconv.String(res.UserInfo.Id) + "-" + gmd5.MustEncryptString(res.UserInfo.UserName) + gmd5.MustEncryptString(res.UserInfo.UserPassword)
	if g.Cfg().MustGet(ctx, "gfToken.multiLogin").Bool() {
		key = gconv.String(res.UserInfo.Id) + "-" + gmd5.MustEncryptString(res.UserInfo.UserName) + gmd5.MustEncryptString(res.UserInfo.UserPassword+ip+userAgent+gtime.Now().String())
	}
	res.UserInfo.UserPassword = ""
	res.Token, err = service.GfToken().GenerateToken(ctx, key, res.UserInfo)
	return
}

func (c *personalController) ResetPwdPersonal(ctx context.Context, req *system.PersonalResetPwdReq) (res *system.PersonalResetPwdRes, err error) {
	res, err = service.Personal().ResetPwdPersonal(ctx, req)
	return
}
