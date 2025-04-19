package service

import (
	"context"

	"fanzhou-go/api/v1/system"
	"fanzhou-go/internal/app/system/dao"
	"fanzhou-go/internal/app/system/model/do"
	"github.com/gogf/gf/v2/frame/g"

	systemConsts "fanzhou-go/internal/app/system/consts"
	"fanzhou-go/library/liberr"
)

type IBigFile interface {
	List(ctx context.Context, req *system.BigFileSearchReq) (res *system.BigFileSearchRes, err error)
	Add(ctx context.Context, req *system.BigFileAddReq, userId uint64) (err error)
	Get(ctx context.Context, id uint64) (res *system.BigFileGetRes, err error)
	Edit(ctx context.Context, req *system.BigFileEditReq, userId uint64) (err error)
	Delete(ctx context.Context, ids []uint64) (err error)
}

func BigFile() IBigFile {
	return new(bigFileTmpl)
}

type bigFileTmpl struct{}

func (s *bigFileTmpl) List(ctx context.Context, req *system.BigFileSearchReq) (res *system.BigFileSearchRes, err error) {
	res = new(system.BigFileSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.BigFile.Ctx(ctx)
		if req != nil {
			if req.Name != "" {
				m = m.Where("name like ?", "%"+req.Name+"%")
			}
		}
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		res.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = systemConsts.PageSize
		}

		var orderBy string = "id asc"
		if req.OrderBy != "" {
			orderBy = req.OrderBy
		}
		err = m.Page(req.PageNum, req.PageSize).Order(orderBy).Scan(&res.List)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}

func (s *bigFileTmpl) Add(ctx context.Context, req *system.BigFileAddReq, userId uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.BigFile.Ctx(ctx).Insert(do.BigFile{
			Name:     req.Name,
			Size:     req.Size,
			Path:     req.Path,
			FullPath: req.FullPath,
			MimeType: req.MimeType,
			Source:   req.Source,
			Describe: req.Describe,
			Md5:      req.Md5,
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *bigFileTmpl) Get(ctx context.Context, id uint64) (res *system.BigFileGetRes, err error) {
	res = new(system.BigFileGetRes)
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.BigFile.Ctx(ctx).WherePri(id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}

func (s *bigFileTmpl) Edit(ctx context.Context, req *system.BigFileEditReq, userId uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.BigFile.Ctx(ctx).WherePri(req.Id).Update(do.BigFile{
			Name:     req.Name,
			Describe: req.Describe,
		})
		liberr.ErrIsNil(ctx, err, "修改错误")
	})
	return
}

func (s *bigFileTmpl) Delete(ctx context.Context, ids []uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.BigFile.Ctx(ctx).WherePri(ids).Delete()
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}
