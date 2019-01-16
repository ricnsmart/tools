package util

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"
)

//MD5单向加密
func Encrypt(password string) (encryptPwd string) {

	salt := time.Now().Unix()

	m5 := md5.New()

	m5.Write([]byte(password))

	m5.Write([]byte(string(salt)))
	//对hash.Hash对象内部存储的数据进行校验和，然后将新的加密内容其追加到data的后面形成一个新的byte切片
	st := m5.Sum(nil)

	encryptPwd = hex.EncodeToString(st)

	return
}

//不同数据的加密值一定不同，相同数据的加密值一定相同，因此，在处理用户登录时只需要对用户传递的密码做加密处理，并与原加密值做比较是否一致即可
func CheckPwd(reqPwd, dataPwd string) error {

	encryptPwd := Encrypt(reqPwd)

	if encryptPwd != dataPwd {
		return errors.New("密码错误")
	}

	return nil
}
