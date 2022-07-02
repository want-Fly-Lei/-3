package service

import (
	"bookSystem/model"
	"bookSystem/utils"
)
//收藏夹服务层

//通过uid查询用户收藏了多少书
func SelectFavoriteByUid(uid int) []model.Favorite {
	var favorites []model.Favorite
	utils.Db.Where("uid = ?", uid).Find(&favorites)
	return favorites
}

//判断是否被用户收藏
func SelectFavoriteByUidAndBid(bid, uid int) model.Favorite  {
	var favorite model.Favorite
	utils.Db.Where("bid = ? and uid = ?",bid,uid).First(&favorite)
	return favorite
}

//添加功能
func AddFavorite(favorite model.Favorite) error {
	res := utils.Db.Create(&favorite)
	return res.Error
}

//删除功能
func DelFavorite(favorite model.Favorite) error {
	res := utils.Db.Where("uid = ? and bid = ?",favorite.Uid, favorite.Cid).Delete(&favorite)
	return res.Error
}