package user

import (
	"context"

	v1 "Capys/api/user/v1"
	"Capys/internal/dao"
	"Capys/internal/model/do"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {

	newUser := do.Users{
		Sid:      req.Sid,
		Username: req.Username,
		Password: req.Password,
	}

	res = &v1.CreateRes{}

	newUserId, err := dao.Users.Ctx(ctx).Data(newUser).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	res.Id = uint64(newUserId)

	return

}
