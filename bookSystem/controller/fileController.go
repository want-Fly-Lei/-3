package controller

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

//保存文件
func SaveFileIntoBooks(c *gin.Context, dst string) error {
	//从请求中读取文件
	var err error
	var myFile *multipart.FileHeader
	myFile, err = c.FormFile("f1")
	if err != nil { //文件读取失败
		return err
	} else {
		err = c.SaveUploadedFile(myFile, dst)
		if err != nil { //文件保存失败
			return err
		} else {
			return nil
		}
	}
}