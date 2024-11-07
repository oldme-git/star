// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	Id        uint        `json:"id"        orm:"id"         ` //
	Username  string      `json:"username"  orm:"username"   ` //
	Password  string      `json:"password"  orm:"password"   ` //
	Email     string      `json:"email"     orm:"email"      ` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
}
