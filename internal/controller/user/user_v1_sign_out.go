package user

import (
	"context"

	"fanzhou-go/api/user/v1"
	"fanzhou-go/internal/service"
)

func (c *ControllerV1) SignOut(ctx context.Context, req *v1.SignOutReq) (res *v1.SignOutRes, err error) {
	err = service.User().SignOut(ctx)
	return
}
