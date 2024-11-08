package users

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"star/internal/dao"
	"star/internal/model"
	"star/internal/model/do"
)

func Register(ctx context.Context, in *model.UsersInput) error {
	if err := CheckUser(ctx, in.Username); err != nil {
		return err
	}

	_, err := dao.Users.Ctx(ctx).Data(do.Users{
		Username: in.Username,
		Password: encryptPassword(in.Password),
		Email:    in.Email,
	}).Insert()
	if err != nil {
		return err
	}
	return nil
}

func CheckUser(ctx context.Context, username string) error {
	count, err := dao.Users.Ctx(ctx).Where("username", username).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("用户已存在")
	}
	return nil
}
