package users

import (
	"context"
	"errors"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
	"star/internal/dao"
	"star/internal/model/entity"
	"star/utility"
)

type userClaims struct {
	Id       uint
	Username string
	jwt.RegisteredClaims
}

func (u *Users) Login(ctx context.Context, username, password string) (tokenString string, err error) {
	var user entity.Users
	err = dao.Users.Ctx(ctx).Where("username", username).Scan(&user)
	if err != nil {
		return "", errors.New("用户名或密码错误")
	}

	if user.Id == 0 {
		return "", errors.New("用户不存在")
	}

	// 将密码加密后与数据库中的密码进行比对
	if user.Password != encryptPassword(password) {
		return "", errors.New("用户名或密码错误")
	}

	// 生成token
	uc := &userClaims{
		Id:       user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	return token.SignedString(utility.JwtKey)
}

func (u *Users) Info(ctx context.Context) (user *entity.Users, err error) {
	user = new(entity.Users)
	tokenString := g.RequestFromCtx(ctx).Request.Header.Get("Authorization")

	tokenClaims, _ := jwt.ParseWithClaims(tokenString, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		return utility.JwtKey, nil
	})

	if claims, ok := tokenClaims.Claims.(*userClaims); ok && tokenClaims.Valid {
		err = dao.Users.Ctx(ctx).Where("id", claims.Id).Scan(&user)
	}
	return
}

func (u *Users) GetUid(ctx context.Context) (uint, error) {
	user, err := u.Info(ctx)
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}
