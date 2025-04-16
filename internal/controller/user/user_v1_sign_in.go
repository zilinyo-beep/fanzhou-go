package user

import (
	"context"

	"fanzhou-go/api/user/v1"
	"fanzhou-go/internal/model"
	"fanzhou-go/internal/service"
)

func (c *ControllerV1) SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error) {
	err = service.User().SignIn(ctx, model.UserSignInInput{
		Passport: req.Passport,
		Password: req.Password,
	})
	return
}
