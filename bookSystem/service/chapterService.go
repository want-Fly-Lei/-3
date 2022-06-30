package service

import (
	"bookSystem/model"
	"bookSystem/utils"
)

//章节服务层
//添加章节
func AddChapter(chapter model.Chapter) error {
	//忽略id字段，其他参数传递过来
	res := utils.Db.Create(&chapter)
	return res.Error
}

//通过查询id获取最后一章
func GetLastChapterByBid(bid int) model.Chapter {
	var chapter model.Chapter
	utils.Db.Where("bid = ?",bid).Last(&chapter)
	return chapter
}

//通过cid和bid查询chapter
func GetChapterContextByBidAndCid(bid, cid int) model.Chapter {
	var chapter model.Chapter
	utils.Db.Where("bid = ? and cid = ?",bid, cid).First(&chapter)
	return chapter
}