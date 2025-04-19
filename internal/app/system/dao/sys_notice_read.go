// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2023-11-09 17:37:25
// 生成路径: internal/app/system/dao/sys_notice_read.go
// 生成人：gfast
// desc:已读记录
// company:云南奇讯科技有限公司
// ==========================================================================

package dao

import (
	"fanzhou-go/internal/app/system/dao/internal"
)

// sysNoticeReadDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type sysNoticeReadDao struct {
	*internal.SysNoticeReadDao
}

var (
	// SysNoticeRead is globally public accessible object for table tools_gen_table operations.
	SysNoticeRead = sysNoticeReadDao{
		internal.NewSysNoticeReadDao(),
	}
)

// Fill with you ideas below.
