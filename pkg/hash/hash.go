package hash

import (
	"gohub/pkg/logger"

	"golang.org/x/crypto/bcrypt"
)

// BcryptHash 使用bcrypt 对密码加密
func BcryptHash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogIf(err)

	return string(bytes)
}

// BcryptCheck 对比明文密码哈希值和数据库里的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func BcryptIsHashed(str string) bool {
	//加密后的长度为60
	return len(str) == 60
}
