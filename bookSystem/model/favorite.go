package model

//收藏夹
type Favorite struct {
	// 用户id
	Uid int `json:"uid" form:"uid"`
	// 书id
	Bid int `json:"bid" form:"bid"`
	//用户上一次读到的章节号
	Cid int `json:"cid" form:"cid"`
}
