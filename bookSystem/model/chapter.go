package model

import "time"

//章节表
type Chapter struct {
	//书号
	Bid     int `json:"bid" form:"bid"`
	// 书章节
	Cid     int `json:"cid" form:"cid"`
	//书内容地址
	Context string `json:"context" form:"context"`
	// 书更新时间
	Time    time.Time `json:"time" form:"time"`
}