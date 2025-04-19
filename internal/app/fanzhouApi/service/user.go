package service

import (
	"errors"
	"fanzhou-go/internal/app/fanzhouApi/dao"
	"fanzhou-go/internal/app/fanzhouApi/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

// 通过手机号查找或注册用户，并返回用户ID
func EnsureUserByPhone(r *ghttp.Request, phone string) (*entity.User, error) {
	ctx := r.GetCtx()

	deviceID := r.Header.Get("CLIENT-UDID")
	deviceModel := r.Header.Get("CLIENT-MODEL")
	clientVersion := r.Header.Get("CLIENT-VERSION")
	loginIP := r.GetClientIp()

	now := gtime.Now()

	user, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Phone, phone).One()
	if err != nil {
		return nil, err
	}

	var userId int64
	var userEntity *entity.User

	if user == nil {
		nickname := "用户" + phone[len(phone)-4:]
		newUser := g.Map{
			"phone":           phone,
			"nickname":        nickname,
			"coin":            5,
			"device_id":       deviceID,
			"device_model":    deviceModel,
			"client_version":  clientVersion,
			"register_ip":     loginIP,
			"last_login_ip":   loginIP,
			"last_login_time": now,
			"created_at":      now,
			"updated_at":      now,
		}

		_, err := dao.User.Ctx(ctx).Insert(newUser)
		if err != nil {
			return nil, err
		}

		user, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Phone, phone).One()
		if err != nil {
			return nil, err
		}
	}
	userId = user["id"].Int64()

	_, _ = dao.User.Ctx(ctx).WherePri(userId).Update(g.Map{
		"last_login_time": now,
		"last_login_ip":   loginIP,
		"device_id":       deviceID,
		"device_model":    deviceModel,
		"client_version":  clientVersion,
		"updated_at":      now,
	})
	_ = user.Struct(&userEntity)
	return userEntity, nil
}
func GetUserById(ctx g.Ctx, userId int64) (*entity.User, error) {
	user, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, userId).One()
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}

	var userEntity *entity.User
	if err := user.Struct(&userEntity); err != nil {
		return nil, err
	}
	return userEntity, nil
}
