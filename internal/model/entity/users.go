// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	Id        uint64      `json:"id"        orm:"id"         description:""` //
	Sid       int64       `json:"sid"       orm:"sid"        description:""` //
	Username  string      `json:"username"  orm:"username"   description:""` //
	Password  string      `json:"-"  orm:"password"   description:""` //
	Money     float64     `json:"money"     orm:"money"      description:""` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""` //
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:""` //
}
