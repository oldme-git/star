package users

import (
	"context"

	"star/internal/dao"
	"star/internal/model/do"
)

func Register(ctx context.Context, username, password, email string) error {
	_, err := dao.Users.Ctx(ctx).Data(do.Users{
		Username: username,
		Password: password,
		Email:    email,
	}).Insert()
	if err != nil {
		return err
	}
	return nil
}
