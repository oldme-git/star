package users

import "github.com/gogf/gf/v2/crypto/gmd5"

func encryptPassword(password string) string {
	return gmd5.MustEncryptString(password)
}
