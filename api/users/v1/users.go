package v1

import "github.com/gogf/gf/v2/frame/g"

type RegisterReq struct {
	g.Meta   `path:"users/register" method:"post" sm:"注册" tags:"用户"`
	Username string `json:"username" v:"required|length:3,12" dc:"用户名"`
	Password string `json:"password" v:"required|length:6,16" dc:"密码"`
	Email    string `json:"email" v:"required|email" dc:"邮箱"`
}

type RegisterRes struct {
}

type LoginReq struct {
	g.Meta   `path:"users/login" method:"post" sm:"登录" tags:"用户"`
	Username string `json:"username" v:"required|length:3,12"`
	Password string `json:"password" v:"required|length:6,16"`
}

type LoginRes struct {
	Token string `json:"token" dc:"在需要鉴权的接口中header加入Authorization: token"`
}
