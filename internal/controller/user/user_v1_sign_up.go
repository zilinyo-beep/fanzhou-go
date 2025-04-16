package user

import (
	"context"

	"fanzhou-go/api/user/v1"
	"fanzhou-go/internal/model"
	"fanzhou-go/internal/service"
)

func (c *ControllerV1) SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error) {
	err = service.User().Create(ctx, model.UserCreateInput{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	return
}
