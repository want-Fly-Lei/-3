package model

type Book struct {
	//书号
	Id int `json:"id" form:"id"`
	//书名
	Bookname string `json:"bookname" form:"bookname"`
	//作者
	Author string `json:"author" form:"author"`
	//介绍
	Description string `json:"description" form:"description"`
	//出版社
	Press string `json:"press" form:"press"`
	//分类
	Kind string `json:"kind" form:"kind"`
	//封面地址
	Cover string `json:"cover" form:"cover"`
}