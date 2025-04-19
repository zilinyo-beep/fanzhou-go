/*
* @desc:登录
* @company:云南奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2022/4/27 21:51
 */

package system

import (
	commonApi "fanzhou-go/api/v1/common"
	"fanzhou-go/internal/app/system/model"
	"github.com/gogf/gf/v2/frame/g"
)

type UserLoginReq struct {
	g.Meta     `path:"/login" tags:"系统后台/登录" method:"post" summary:"用户登录"`
	Username   string `p:"username" v:"required#用户名不能为空"`
	Password   string `p:"password" v:"required#密码不能为空"`
	VerifyCode string `p:"verifyCode"`
	VerifyKey  string `p:"verifyKey"`
}

type UserLoginRes struct {
	g.Meta      `mime:"application/json"`
	UserInfo    *model.LoginUserRes `json:"userInfo"`
	Token       string              `json:"token"`
	MenuList    []*model.UserMenus  `json:"menuList"`
	Permissions []string            `json:"permissions"`
}

type UserLoginOutReq struct {
	g.Meta `path:"/logout" tags:"系统后台/登录" method:"get" summary:"退出登录"`
	commonApi.Author
}

type UserLoginOutRes struct {
}
