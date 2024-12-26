package users

import "github.com/gogf/gf/v2/crypto/gmd5"

type Users struct {
}

func New() *Users {
	return &Users{}
}

func (u *Users) encryptPassword(password string) string {
	return gmd5.MustEncryptString(password)
}
