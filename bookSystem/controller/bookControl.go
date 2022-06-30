package controller

import (
	"fmt"
	"net/http"

	"bookSystem/model"
	"bookSystem/service"
	"bookSystem/utils"

	"github.com/gin-gonic/gin"
)

//书本控制层

//添加书本
func AddBook(ctx *gin.Context) {
	var book model.Book
	var err error 
	if err = ctx.ShouldBind(&book); err == nil {
		var id int = service.GetLastBook().Id
		book.Cover = fmt.Sprintf("%s/books/%d/%s",utils.MY_PROJECT_FILE_PATH,id + 1, utils.COVER_FILE_NAME)
		if err = service.AddBook(book); err == nil { //插入数据库成功
			if err = SaveFileIntoBooks(ctx,book.Cover); err == nil { //写入本地文件成功
				ctx.JSON(http.StatusOK, gin.H { //添加成功
				"msg":"ok",
			})
			} else { //写入封面失败
				ctx.JSON(http.StatusUnprocessableEntity, gin.H {
					"msg":err.Error(),
				})
			}
			
		} else {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H {
				"msg":err.Error(),
			})
		}
		
	} else { //绑定失败
		ctx.JSON(http.StatusBadRequest, gin.H {
			"msg":err.Error(),
		})
	}
}

//所有书本
func AllBook(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.SelectAllBook())
}

//同分类查询书
func SelectBookByKind(ctx *gin.Context) {
	var kind string = ctx.Query("kind")
	fmt.Println(kind)
	var books []model.Book = service.GetBooksByKind(kind)
	if len(books) == 0 || books == nil {
		ctx.JSON(http.StatusBadRequest, gin.H {
			"msg":"该分类暂无书籍",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H {
			"msg":"",
			"books":books,
		})
	}
}