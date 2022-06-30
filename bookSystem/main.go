package main

import (
	"bookSystem/controller"

	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	//user下的接口
	userGroup := r.Group("/user")
	{
		//用户登陆
		userGroup.POST("/login",controller.UserLogin)
		//用户注册
		userGroup.POST("/register", controller.UserRegister)
		//修改用户
		userGroup.POST("/reset", controller.ReSetUser)
	}

	//book下的接口
	bookGroup := r.Group("/book")
	{
		//查询所有书籍
		bookGroup.GET("/allBook", controller.AllBook)
		//添加书籍
		bookGroup.POST("/addBook", controller.AddBook)
		//通过分类返回图书信息
		bookGroup.GET("/selectBookByKind",controller.SelectBookByKind)
	}

	chapterGroup := r.Group("/chapter") 
	{
		//新增章节
		chapterGroup.POST("/addChapter",controller.AddChapter)
		//查询章节内容
		chapterGroup.GET("/showChapterContext",controller.ShowChapterContext)
	}

	r.Run(":8086")
}
