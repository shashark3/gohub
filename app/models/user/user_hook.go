package user

import (
	"gohub/pkg/hash"

	"gorm.io/gorm"
)

// BeforeSave  Gorm的模型钩子，用于在模型创建和更新之前调用  利用此机制对密码进行加密
func (userModel *User) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(userModel.Password) {
		userModel.Password = hash.BcryptHash(userModel.Password)
	}
	return
}
