/*
* @desc:在线用户
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2023/1/10 16:57
 */

package system

import (
	commonApi "fanzhou-go/api/v1/common"
	"fanzhou-go/internal/app/system/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserOnlineSearchReq 列表搜索参数
type SysUserOnlineSearchReq struct {
	g.Meta   `path:"/online/list" tags:"系统后台/在线用户管理" method:"get" summary:"列表"`
	Username string `p:"userName"`
	Ip       string `p:"ipaddr"`
	commonApi.PageReq
	commonApi.Author
}

// SysUserOnlineSearchRes 列表结果
type SysUserOnlineSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*entity.SysUserOnline `json:"list"`
}

type SysUserOnlineForceLogoutReq struct {
	g.Meta `path:"/online/forceLogout" tags:"系统后台/在线用户管理" method:"delete" summary:"强制用户退出登录"`
	commonApi.Author
	Ids []int `p:"ids" v:"required#ids不能为空"`
}

type SysUserOnlineForceLogoutRes struct {
	commonApi.EmptyRes
}
