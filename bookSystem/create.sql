CREATE DATABASE  IF not EXISTS booksystem;
USE booksystem;

# 用户表
CREATE TABLE `users` (
	id int PRIMARY KEY auto_increment,
	username VARCHAR(100) NOT NULL,
	`password` VARCHAR(15) NOT NULL,
	email VARCHAR(20),
	root INT(1)
);

# 书表
CREATE TABLE books (
id int PRIMARY KEY auto_increment,
bookname VARCHAR(40) not NULL,
author VARCHAR(40),
description VARCHAR(200),
press VARCHAR(40), # 出版社
kind VARCHAR(20), # 分类
cover VARCHAR(100) # 图书封面地址
);

# 章节表
CREATE TABLE chapters (
	bid int, # 书名
	cid int, # 章节号
	context  VARCHAR(100), # 章节存放的地址
	time TIMESTAMP, # 章节更新修改时间
	PRIMARY KEY(bid,cid)  # 书id和章节id同时作为主码
);

# 收藏夹
CREATE TABLE favorites (
	uid int, # 用户id
	bid int, # 书id
	cid int # 用户上次最后读取的章节
);
