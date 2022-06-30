package service

import (
	"bookSystem/model"
	"bookSystem/utils"
)

//章节服务层
//添加章节
func AddChapter(chapter model.Chapter) error {
	var err error
	//忽略id字段，其他参数传递过来
	res := utils.Db.Create(&chapter)
	if res.Error != nil {
		return res.Error
	}
	return err
}

//通过查询id获取最后一章
func GetLastChapterByBid(bid int) model.Chapter {
	var chapter model.Chapter
	utils.Db.Find("bid = ?",bid).Last(&chapter)
	return chapter
}