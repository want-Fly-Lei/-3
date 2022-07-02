package model

type User struct {
	//用户名
	Username string `gorm:"varchar(100);not null;unique" form:"username" json:"username"`
	//邮箱
	Email string `gorm:"varchar(100)" form:"email" json:"email"`
	//密码
	Password string `gorm:"varchar(100);not null" form:"password" json:"password"`
	//id
	Id int `gorm:"primary_key" form:"id" json:"id"`
	//权限，0用普通用户，1为管理员
	Root int `gorm:"int(1)" form:"root" json:"root"`
}
