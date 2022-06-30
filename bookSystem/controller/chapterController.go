package controller

import (
	"bookSystem/model"
	"bookSystem/service"
	"bookSystem/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//添加章节
func AddChapter(ctx *gin.Context) {
	var chapter model.Chapter
	chapter.Time = time.Now()
	var err error
	if err = ctx.ShouldBind(&chapter); err == nil {
		chapter.Cid = service.GetLastChapterByBid(chapter.Bid).Cid + 1
		//传入地址
		chapter.Context = fmt.Sprintf("%s/books/%d/%d.txt",utils.MY_PROJECT_FILE_PATH,chapter.Bid,chapter.Cid)
		if err = service.AddChapter(chapter); err == nil { //添加数据成功
			if err = SaveFileIntoBooks(ctx,chapter.Context); err == nil { //写入文件成功
				ctx.JSON(http.StatusOK, gin.H {
					"msg":"ok",
				})
			} else {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H {
					"msg":err.Error(),
				})
			}
		} else {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H {
				"msg":err.Error(),
			})
		}
	} else { // 绑定失败
		ctx.JSON(http.StatusUnprocessableEntity, gin.H {
			"msg":err.Error(),
		})
	}
}