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
	var err error
	if err = ctx.ShouldBind(&chapter); err == nil {
		//fmt.Println(chapter)
		chapter.Cid = service.GetLastChapterByBid(chapter.Bid).Cid + 1
		//传入地址
		chapter.Context = fmt.Sprintf("%s/books/%d/%d.txt", utils.MY_PROJECT_FILE_PATH, chapter.Bid, chapter.Cid)
		chapter.Time = time.Now()
		if err = service.AddChapter(chapter); err == nil { //添加数据成功
			if err = SaveFileIntoBooks(ctx, chapter.Context); err == nil { //写入文件成功
				ctx.JSON(http.StatusOK, gin.H{
					"msg": "ok",
				})
			} else {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"msg": err.Error(),
				})
			}
		} else {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"msg": err.Error(),
			})
		}
	} else { // 绑定失败
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": err.Error(),
		})
	}
}

//显示章节内容
func ShowChapterContext(ctx *gin.Context) {
	var chapter model.Chapter
	var err error
	if err = ctx.ShouldBind(&chapter); err == nil {
		var str string
		if chapter = service.GetChapterContextByBidAndCid(chapter.Bid, chapter.Cid); chapter.Context != "" { //地址非空可以查询
			if str = utils.FileToString(chapter.Context); str != "" { //写入文件成功
				ctx.JSON(http.StatusOK, gin.H{
					"msg": "",
					"context":str,
				})
			} else {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"msg": "内容为空",
				})
			}
		} else {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"msg": "无法查询",
			})
		}
	} else { // 绑定失败
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": err.Error(),
		})
	}
}

//更新指定章节章节内容
func UpdateChapterContext(ctx *gin.Context) {
	var chapter model.Chapter
	var err error
	if err = ctx.ShouldBind(&chapter); err == nil {
		//fmt.Println(chapter)
		if chapter.Cid >= 0 && chapter.Cid <= service.GetLastChapterByBid(chapter.Bid).Cid {
			//说明传入章节数不对
			ctx.JSON(http.StatusBadRequest, gin.H {
				"msg":"你传入的章节数不对",
			})
			return 
		}
		//传入地址
		chapter.Context = fmt.Sprintf("%s/books/%d/%d.txt", utils.MY_PROJECT_FILE_PATH, chapter.Bid, chapter.Cid)
		chapter.Time = time.Now()
		if err = service.UploadChapterByCid(chapter); err == nil { //更新数据成功
			if err = SaveFileIntoBooks(ctx, chapter.Context); err == nil { //写入文件成功
				ctx.JSON(http.StatusOK, gin.H{
					"msg": "ok",
				})
			} else {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"msg": err.Error(),
				})
			}
		} else {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"msg": err.Error(),
			})
		}
	} else { // 绑定失败
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": err.Error(),
		})
	}
}