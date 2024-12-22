package account

import (
	"context"

	"star/api/account/v1"
)

func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
	user, err := c.users.Info(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.InfoRes{
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
