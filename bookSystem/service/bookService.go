package service

import (
	"bookSystem/model"
	"bookSystem/utils"
)

//返回数据库服务book函数

//添加书籍
func AddBook(book model.Book) error {
	var err error
	//忽略id字段，其他参数传递过来
	res := utils.Db.Omit("id").Create(&book)
	if res.Error != nil {
		return res.Error
	}
	return err
}

//查询所有书籍
func SelectAllBook() []model.Book {
	var books []model.Book
	utils.Db.Find(&books)
	return books
}

//获取最后一本书
func GetLastBook() model.Book {
	var book model.Book
	utils.Db.Last(&book)
	return book
}

//通过分类查询书籍
func GetBooksByKind(kind string) []model.Book{
	var books []model.Book
	utils.Db.Where("kind = ?", kind).Find(&books)
	return books
}

//通过模糊查询书籍
func GetBooksByBookNameMohu(bookname string) []model.Book{
	var books []model.Book
	utils.Db.Where("bookname like ?", utils.ToMOhu(bookname)).Find(&books)
	return books
}

//通过书id查找
func GetBookByBid(bid int) model.Book{
	var book model.Book
	utils.Db.First(&book, bid)
	return book
}

//通过书id更新图书信息
func UpdateBookById(id int, book model.Book) error{
	res := utils.Db.Update(&book)
	return res.Error
}