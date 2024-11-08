package v1

import "github.com/gogf/gf/v2/frame/g"

type RegisterReq struct {
	g.Meta   `path:"users/register" method:"post"`
	Username string `json:"username" v:"required|length:3,12"`
	Password string `json:"password" v:"required|length:6,16"`
	Email    string `json:"email" v:"required|email"`
}

type RegisterRes struct {
}
