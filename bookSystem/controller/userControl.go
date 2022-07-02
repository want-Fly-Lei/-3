package controller

import (
	"bookSystem/model"
	"bookSystem/service"
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//用户控制层

//用户登陆
func UserLogin(ctx *gin.Context) {
	var err error
	var user model.User
	//把传入的参数绑定到user
	if err = ctx.ShouldBind(&user); err == nil {
		//未找到用户
		if service.SelectUserByUserName(user.Username).Id == 0 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"msg": "该用户不存在",
			})
		} else {
			if user = service.SelectUserLogin(user.Username, user.Password); user.Id == 0 {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"msg": "账号或者密码错误",
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"msg":      "",
					"root":     user.Root,     //返回用户权限
					"id":       user.Id,       //返回用户id
					"username": user.Username, //返回用户名
				})
			}
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}
}

//用户注册
func UserRegister(ctx *gin.Context) {
	var err error
	var user model.User
	//把传入的参数绑定到user
	if err = ctx.ShouldBind(&user); err == nil {
		//未找到用户
		if service.SelectUserByUserName(user.Username).Id == 0 {
			if err = service.RegisterUser(user); err == nil {
				ctx.JSON(http.StatusOK, gin.H{
					"msg": "ok",
				})
			} else {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"msg": err,
				})
			}
		} else {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"msg": "该用户已存在",
			})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}
}

//用户修改功能
func ReSetUser(ctx *gin.Context) {
	var err error
	var user model.User
	//把传入的参数绑定到user
	if err = ctx.ShouldBind(&user); err == nil {
		user.Root = 0 // 防止乱搞
		if err = service.ResetUserById(user.Id,user); err == nil {
			ctx.JSON(http.StatusOK, gin.H {
				"msg":"ok",
			})
		} else {
			ctx.JSON(http.StatusRequestEntityTooLarge, gin.H {
				"msg":err.Error(),
			})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}
}

//管理员修改功能
func ReSetUserAdmin(ctx *gin.Context) {
	var err error
	var user model.User
	//把传入的参数绑定到user
	if err = ctx.ShouldBind(&user); err == nil {
		//fmt.Println(user)
		if err = service.ResetUserById(user.Id,user); err == nil {
			ctx.JSON(http.StatusOK, gin.H {
				"msg":"ok",
			})
		} else {
			ctx.JSON(http.StatusRequestEntityTooLarge, gin.H {
				"msg":err.Error(),
			})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}
}

//所搜所有用户
func AllUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.SelectAllUse())
}