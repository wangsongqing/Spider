package user

import (
	"Spider/pkg/database"
)

// IsEmailExist 判断 Email 是否已经存在
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断 手机号 是否已经存在
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

func GetUserById(id int64) (UserModel User) {
	database.DB.Model(User{}).Where("id = ?", id).Find(&UserModel)
	return
}
