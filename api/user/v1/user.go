package v1

import (
	"Capys/internal/model/entity"
	"Capys/utility/enum"

	"github.com/gogf/gf/v2/frame/g"
)

// 创建用户的请求
type CreateReq struct {
	g.Meta   `path:"/user" method:"post" tags:"用户" summary:"创建用户" `
	Sid      string `v:"required|length:5,5" dc:"学号" json:"sid"`
	Username string `v:"required|length:1,10" dc:"用户名" json:"username"`
	Password string `v:"required|length:6,16" dc:"密码" json:"password"`
}

// 创建用户的响应体
type CreateRes struct {
	Id uint64 `json:"id" dc:"用户id"`
}

type DeleteReq struct {
	g.Meta `path:"/user/{id}" method:"delete" tags:"用户" summary:"删除用户"`
	Id     uint64 `v:"required" dc:"密码"`
}

type DeleteRes struct{}

type UpdateReq struct {
	g.Meta   `path:"/user/{id}" method:"put" tags:"用户" summary:"修改用户"`
	Username *string      `v:"required|length:1,10" dc:"用户名"`
	Password *string      `v:"required|length:6,16" dc:"密码"`
	Status   *enum.Status `v:"in:0,1" dc:"密码"`
}
type UpdateRes struct{}

type GetOneReq struct {
	g.Meta `path:"/user/{id}" method:"get" tags:"用户" summary:"查询一个用户"`
	Id     uint64 `v:"required" dc:"用户id"`
}

type GetOneRes struct {
	*entity.Users `dc:"用户"`
}

type GetListReq struct {
	g.Meta `path:"/user" method:"get" tags:"用户" summary:"Get users"`
	Money  float64      `dc:"user age"`
	Status *enum.Status `v:"in:0,1" dc:"user status"`
}
type GetListRes struct {
	List []*entity.Users `json:"list" dc:"user list"`
}
