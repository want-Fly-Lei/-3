package controller

import (
	"bookSystem/model"
	"bookSystem/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//添加收藏夹
func AddFavorite(ctx *gin.Context) {
	var favorite model.Favorite
	var err error
	//把传入的参数绑定到user
	if err = ctx.ShouldBind(&favorite); err == nil {
		favorite.Cid = 1 //添加默认是第一章
		if service.GetChapterContextByBidAndCid(favorite.Bid,favorite.Cid).Bid == 0 {
			if err = service.AddFavorite(favorite); err == nil {
				ctx.JSON(http.StatusOK, gin.H{
					"msg": "ok",
				})
			} else {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"msg": err,
				})
			}
		} else {//未找到用户
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"msg": "该用户以把书添加到喜欢",
			})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}
}

//删除收藏夹
func DeleteFavorite(ctx *gin.Context) {
	var favorite model.Favorite
	var err error
	//把传入的参数绑定到user
	if err = ctx.ShouldBind(&favorite); err == nil {
		favorite.Cid = 1 //添加默认是第一章
		if service.GetChapterContextByBidAndCid(favorite.Bid,favorite.Cid).Bid != 0 {
			if err = service.DelFavorite(favorite); err == nil {
				ctx.JSON(http.StatusOK, gin.H{
					"msg": "ok",
				})
			} else {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"msg": err,
				})
			}
		} else {//未找到用户
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"msg": "该用户没收藏对应",
			})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}
}
