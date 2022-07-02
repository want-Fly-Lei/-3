# 1.运行
## 1.1go语言环境
环境安装这里不赘述

## 1.2 go mod
该项目是使用go mod方式创建，因为使用了gin和gorm框架，所以需要下载依赖
```go
go mod tidy
```

## 1.3运行
在本文件夹下执行
```go
go run main.go
```

# 2.框架介绍
## 2.1gin框架
Gin 是一个用 Go (Golang) 编写的 Web 框架。 它具有类似 martini 的 API，性能要好得多，多亏了 httprouter，速度提高了 40 倍。 如果您需要性能和良好的生产力，您一定会喜欢 Gin。[官方文档](https://gin-gonic.com/zh-cn/docs/introduction/).

## 2.2gorm框架
GORM是一个使用Go语言编写的ORM框架。它文档齐全，对开发者友好，支持主流数据库。[推荐中文文档](https://www.w3cschool.cn/gormdoc/gormdoc-k3o23ltb.html).

# 3.各个文件
model数数据层，service是数据库的各种操作，controller是json的api接口,utils是自定义的工具类，books是存放图书信息的地方(书的内容，书的封面)

# 4.其他
我在utils中的config.go我觉得每次都读取项目绝对地址太麻烦了，所以地址写死了，自己进行修改即可