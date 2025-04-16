package user

import (
	"context"

	"fanzhou-go/api/user/v1"
	"fanzhou-go/internal/service"
)

func (c *ControllerV1) Profile(ctx context.Context, req *v1.ProfileReq) (res *v1.ProfileRes, err error) {
	res = &v1.ProfileRes{
		User: service.User().GetProfile(ctx),
	}
	return
}
