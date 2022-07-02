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
		userGroup.POST("/login", controller.UserLogin)
		//用户注册
		userGroup.POST("/register", controller.UserRegister)
		//修改用户
		userGroup.POST("/reset", controller.ReSetUser)
		//管理员修改用户
		userGroup.POST("/resetAdmin", controller.ReSetUserAdmin)
		//搜索所有用户
		userGroup.GET("/allUser", controller.AllUser)
	}

	//book下的接口
	bookGroup := r.Group("/book")
	{
		//查询所有书籍
		bookGroup.GET("/allBook", controller.AllBook)
		//添加书籍
		bookGroup.POST("/addBook", controller.AddBook)
		//通过分类返回图书信息
		bookGroup.GET("/selectBookByKind", controller.SelectBookByKind)
		//书名模糊查询
		bookGroup.GET("/selectByBookNameMoHu", controller.SelectByBookNameMoHu)
		//通过书id查询
		bookGroup.GET("/selectBookByBid", controller.SelectBookByBid)
		//更细书的信息
		bookGroup.POST("/updateBook", controller.UpdateBook)
	}

	chapterGroup := r.Group("/chapter")
	{
		//新增章节
		chapterGroup.POST("/addChapter", controller.AddChapter)
		//查询章节内容
		chapterGroup.GET("/showChapterContext", controller.ShowChapterContext)
		//跟新章节内容
		chapterGroup.POST("/updateChapterContext", controller.UpdateChapterContext)
	}

	favoriteGroup := r.Group("/favorite")
	{
		//添加收藏夹
		favoriteGroup.GET("/addFavorite", controller.AddFavorite)
		//删除收藏夹
		favoriteGroup.GET("/deleteFavorite", controller.DeleteFavorite)
	}

	r.Run(":8086")
}
