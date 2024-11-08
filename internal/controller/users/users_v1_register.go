package users

import (
	"context"

	"star/internal/logic/users"
	"star/internal/model"

	"star/api/users/v1"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	err = users.Register(ctx, &model.UsersInput{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})
	return nil, err
}
