package user

import (
	"context"

	v1 "Capys/api/user/v1"
	"Capys/internal/dao"
)

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	res = &v1.GetOneRes{}

	dao.Users.Ctx(ctx).WherePri(req.Id).Scan(&res.Users)
	return
}
