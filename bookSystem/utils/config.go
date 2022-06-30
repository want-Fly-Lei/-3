package utils

//存储项目相关的配置代码

var (
	//文件工程路径
	MY_PROJECT_FILE_PATH string = "E:/vscode_workplace/go/src/course3/bookSystem"
	COVER_FILE_NAME string = "icon.jpg" //图片存储名称
)

//模糊查询
func toMOhu(condition string) string {
	return "%" + condition + "%"
}