package service

import (
	"bookSystem/utils"
	"bookSystem/model"
)

//返回数据库用户数据

func SelectUserByUserName(username string) model.User {
	var user model.User
	utils.Db.Where("username = ?", username).First(&user)
	return user
}

func SelectUserLogin(username, password string) model.User {
	var user model.User
	utils.Db.Where("username = ? and password = ?", username, password).First(&user)
	return user
}

func RegisterUser(user model.User) error {
	user.Root = 0 //防止权限问题
	res := utils.Db.Select("Username", "Email", "Password", "Root").Create(&user)
	if res.Error != nil {
		return res.Error
	}
	
	return nil
}