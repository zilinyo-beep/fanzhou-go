package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"fanzhou-go/api/user/v1"
	"fanzhou-go/internal/service"
)

func (c *ControllerV1) CheckPassport(ctx context.Context, req *v1.CheckPassportReq) (res *v1.CheckPassportRes, err error) {
	available, err := service.User().IsPassportAvailable(ctx, req.Passport)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf(`Passport "%s" is already token by others`, req.Passport)
	}
	return
}
