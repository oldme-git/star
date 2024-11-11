package account

import (
	"context"

	"star/api/account/v1"
	"star/internal/logic/users"
)

func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
	user, err := users.Info(ctx)
	if err == nil {
		res = &v1.InfoRes{
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}
	return
}
