// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2023-11-09 11:41:17
// 生成路径: internal/app/system/logic/sys_notice.go
// 生成人：gfast
// desc:通知公告
// company:云南奇讯科技有限公司
// ==========================================================================

package logic

import (
	"context"
	"database/sql"
	"errors"
	"fanzhou-go/internal/app/system/consts"
	"fanzhou-go/internal/app/system/dao"
	sysDao "fanzhou-go/internal/app/system/dao"
	"fanzhou-go/internal/app/system/model"
	"fanzhou-go/internal/app/system/model/do"
	"fanzhou-go/internal/app/system/service"
	"fanzhou-go/library/libWebsocket"
	"fanzhou-go/library/liberr"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func init() {
	service.RegisterSysNotice(New())
}

func New() service.ISysNotice {
	return &sSysNotice{}
}

type sSysNotice struct{}

func (s *sSysNotice) List(ctx context.Context, req *model.SysNoticeSearchReq) (listRes *model.SysNoticeSearchRes, err error) {
	listRes = new(model.SysNoticeSearchRes)
	currentUserId := service.Context().GetUserId(ctx)
	if currentUserId <= 0 {
		err = errors.New("用户信息查询失败")
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysNotice.Ctx(ctx).WithAll().As("n")
		if req.Id != "" {
			m = m.Where("n."+dao.SysNotice.Columns().Id+" = ?", req.Id)
		}
		if req.Title != "" {
			m = m.Where("n."+dao.SysNotice.Columns().Title+" = ?", req.Title)
		}
		if req.Type != "" {
			m = m.Where("n."+dao.SysNotice.Columns().Type+" = ?", gconv.Int64(req.Type))
		}
		if req.Tag != "" {
			m = m.Where("n."+dao.SysNotice.Columns().Tag+" = ?", gconv.Int(req.Tag))
		}
		if req.Status != "" {
			m = m.Where("n."+dao.SysNotice.Columns().Status+" = ?", gconv.Int(req.Status))
		}
		if len(req.DateRange) != 0 {
			m = m.Where("n."+dao.SysNotice.Columns().CreatedAt+" >=? AND "+dao.SysNotice.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
		}
		listRes.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取总行数失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		listRes.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		order := "n.id desc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		m = m.LeftJoin("sys_notice_read as nr", "nr.notice_id=n.id")
		//Where("nr.user_id=?", currentUserId)
		var res []*model.SysNoticeListRes

		err = m.Page(req.PageNum, req.PageSize).Fields("" +
			"n.*," +
			"SUM(nr.clicks) as clickNumber" +
			//"(nr.user_id=" + strconv.FormatUint(currentUserId, 10) + ") as isRead" +
			"").Order(order).Group("n.id").Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = res

	})
	return
}
func (s *sSysNotice) ListShow(ctx context.Context, req *model.SysNoticeSearchReq) (listRes *model.SysNoticeSearchRes, err error) {
	listRes = new(model.SysNoticeSearchRes)
	currentUserId := service.Context().GetUserId(ctx)
	if currentUserId <= 0 {
		err = errors.New("用户信息查询失败")
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysNotice.Ctx(ctx).WithAll().As("n")
		m = m.LeftJoin("sys_notice_read as nr", fmt.Sprintf("nr.notice_id=n.id AND nr.user_id=%d", currentUserId))
		if req.Id != "" {
			m = m.Where("n."+dao.SysNotice.Columns().Id+" = ?", req.Id)
		}
		if req.Title != "" {
			m = m.Where("n."+dao.SysNotice.Columns().Title+" = ?", req.Title)
		}
		if req.Type != "" {
			m = m.Where("n."+dao.SysNotice.Columns().Type+" = ?", gconv.Int64(req.Type))
			if gconv.Int(req.Type) == consts.SysLetterType {
				if service.ToolsGenTable().IsMysql() {
					m = m.Where(fmt.Sprintf("JSON_CONTAINS(n.receiver,'%d')", currentUserId))
				} else {
					m = m.Where(fmt.Sprintf("receiver::jsonb @> '%d'::jsonb", currentUserId))
				}
			}
		}
		if req.Tag != "" {
			m = m.Where("n."+dao.SysNotice.Columns().Tag+" = ?", gconv.Int(req.Tag))
		}
		if req.Status != "" {
			m = m.Where("n."+dao.SysNotice.Columns().Status+" = ?", gconv.Int(req.Status))
		}
		if len(req.DateRange) != 0 {
			m = m.Where("n."+dao.SysNotice.Columns().CreatedAt+" >=? AND "+dao.SysNotice.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
		}
		listRes.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取总行数失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		listRes.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		order := "n.id desc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		var res []*model.SysNoticeListRes
		err = m.Page(req.PageNum, req.PageSize).Fields("" +
			"n.*,nr.id IS NOT NULL as isRead").Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = res
	})
	return
}
func (s *sSysNotice) GetById(ctx context.Context, id int64) (res *model.SysNoticeInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysNotice.Ctx(ctx).WithAll().Where(dao.SysNotice.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sSysNotice) Add(ctx context.Context, req *model.SysNoticeAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		data := do.SysNotice{
			Title:     req.Title,
			Type:      req.Type,
			Tag:       req.Tag,
			Content:   req.Content,
			Remark:    req.Remark,
			Receiver:  req.Receiver,
			Sort:      req.Sort,
			Status:    req.Status,
			CreatedBy: req.CreatedBy,
		}
		var rows sql.Result
		rows, err = dao.SysNotice.Ctx(ctx).Insert(data)
		liberr.ErrIsNil(ctx, err, "添加失败")
		data.Id, err = rows.LastInsertId()
		liberr.ErrIsNil(ctx, err, "获取ID失败")
		response := &libWebsocket.WResponse{
			Event: "notice",
			Data:  data,
		}
		if req.Type == consts.SysLetterType {
			//系统私信
			if len(req.Receiver) > 0 {
				for _, id := range req.Receiver {
					libWebsocket.SendToUser(id, response)
				}
			}
		} else {
			//系统通知
			libWebsocket.SendToAll(response)
		}
	})
	return
}

func (s *sSysNotice) Edit(ctx context.Context, req *model.SysNoticeEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		data := do.SysNotice{
			Title:     req.Title,
			Tag:       req.Tag,
			Content:   req.Content,
			Remark:    req.Remark,
			Sort:      req.Sort,
			Status:    req.Status,
			Receiver:  req.Receiver,
			UpdatedBy: req.UpdatedBy,
		}
		_, err = dao.SysNotice.Ctx(ctx).WherePri(req.Id).Update(data)
		liberr.ErrIsNil(ctx, err, "修改失败")
		data.Id = req.Id
		data.Type = req.Type
		response := &libWebsocket.WResponse{
			Event: "notice",
			Data:  data,
		}
		if req.Type == consts.SysLetterType {
			//系统私信
			if len(req.Receiver) > 0 {
				for _, id := range req.Receiver {
					libWebsocket.SendToUser(id, response)
				}
			}
		} else {
			//系统通知
			libWebsocket.SendToAll(response)
		}
	})
	return
}

func (s *sSysNotice) Delete(ctx context.Context, ids []int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysNotice.Ctx(ctx).Delete(dao.SysNotice.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

/*
	func (s *sSysNotice) IndexData(ctx context.Context) (res *model.SysNoticeIndexRes, err error) {
		err = g.Try(ctx, func(ctx context.Context) {
			userInfo := service.Context().GetLoginUser(ctx)
			readIdMap, err := dao.SysNoticeRead.Ctx(ctx).Where(dao.SysNoticeRead.Columns().UserId, userInfo.Id).Fields(dao.SysNoticeRead.Columns().NoticeId).All()
			fmt.Println(readIdMap, err)
			dao.SysNotice.Ctx(ctx).Where(dao.SysNotice.Columns().Type, "1").Order("id desc").Scan(res.Type1List)
			dao.SysNotice.Ctx(ctx)

		})
		return
	}
*/
//未读消息列表
func (s *sSysNotice) UnReadList(ctx context.Context) (res *model.SysNoticeListRes, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *sSysNotice) UnReadCount(ctx context.Context, currentUserId uint64) (sysNoticeUnreadCount *model.SysNoticeUnreadCount, err error) {
	if currentUserId <= 0 {
		err = errors.New("获取用户信息失败")
		return
	}
	countFunc := func(t int) (count int) {
		noticeIds, serr := s.CurrentUseWithIds(ctx, currentUserId, t)
		if serr != nil {
			return 0
		}
		all, serr := dao.SysNotice.Ctx(ctx).
			WhereIn(dao.SysNotice.Columns().Id, noticeIds).
			Count()
		if serr != nil {
			return 0
		}
		read, serr := dao.SysNoticeRead.Ctx(ctx).
			WhereIn(dao.SysNoticeRead.Columns().NoticeId, noticeIds).
			Where(dao.SysNoticeRead.Columns().UserId, currentUserId).
			Count()
		if serr != nil {
			return 0
		}
		count = all - read
		return
	}
	sysNoticeUnreadCount = new(model.SysNoticeUnreadCount)
	sysNoticeUnreadCount.NotifyCount = countFunc(consts.SysNoticeType) //获取未读通知数量
	sysNoticeUnreadCount.NoticeCount = countFunc(consts.SysLetterType) //获取未读私信数量
	return
}

func (s *sSysNotice) ReadAll(ctx context.Context, nType string) (err error) {
	currentUserId := service.Context().GetUserId(ctx)
	if currentUserId <= 0 {
		err = errors.New("获取当前用户信息失败")
		return
	}
	subQuery := g.Model("sys_notice_read").Where("user_id=?", currentUserId).Fields("id")
	/*	dao.SysNoticeRead.Ctx(ctx).Where(dao.SysNoticeRead.Columns().UserId, currentUserId)*/
	unReadIdMap, err := dao.SysNotice.Ctx(ctx).WhereNotIn("id", subQuery).Fields("id").All()
	if unReadIdMap != nil {
		insertList := make([]*model.SysNoticeReadAddReq, 0)
		for _, record := range unReadIdMap {
			insertList = append(insertList, &model.SysNoticeReadAddReq{
				NoticeId:  record["id"].Int64(),
				UserId:    int64(currentUserId),
				CreatedAt: gtime.Now(),
			})
		}
		_, err = dao.SysNoticeRead.Ctx(ctx).Data(insertList).Insert()
		if err != nil {
			return
		}
	}
	return
}

func (s *sSysNotice) GetUserNameList(ctx context.Context, search string) (res []*model.SysNoticeUserNickname, err error) {
	res = make([]*model.SysNoticeUserNickname, 0)
	err = sysDao.SysUser.Ctx(ctx).Where("user_nickname like ?", "%"+search+"%").Fields("id,user_nickname").Page(1, 6).Scan(&res)
	return
}

func (s *sSysNotice) NoticeReadAddUserId(ctx context.Context, req *model.SysNoticeReadAddUserReq) (err error) {

	return
}

func (s *sSysNotice) CurrentUseWithIds(ctx context.Context, currentUserId uint64, noticeType int) (ids []int64, err error) {
	m := dao.SysNotice.Ctx(ctx)
	m = m.Where(dao.SysNotice.Columns().Status, 1).
		Where(dao.SysNotice.Columns().Type, noticeType)
	if noticeType == consts.SysLetterType {
		if service.ToolsGenTable().IsMysql() {
			m = m.Where("JSON_CONTAINS(`receiver`,'" + fmt.Sprintf("%d", currentUserId) + "')")
		} else {
			m = m.Where("receiver::jsonb @> '" + gconv.String(currentUserId) + "'::jsonb")
		}
	}
	all, err := m.Fields("id").All()
	if err != nil {
		return
	}
	for _, record := range all {
		ids = append(ids, record["id"].Int64())
	}
	return
}
