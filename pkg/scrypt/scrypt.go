package scrypt

import (
	"golang.org/x/crypto/bcrypt"
)

// HashAndSalt 注册密码入库加密
func HashAndSalt(pwdStr string) (pwdHash string, err error) {
	pwd := []byte(pwdStr)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return
	}
	pwdHash = string(hash)
	return
}

// ComparePasswords 登录解密,
func ComparePasswords(hashedPwd string, plainPwd string) (err error) {
	byteHash := []byte(hashedPwd)
	bytePwd := []byte(plainPwd)
	err = bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	return err
}
