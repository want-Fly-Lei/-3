package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func FileToString(dst string) string {
	var myFile *os.File
	var err error
	//os.O_RDWR表示可读可写
	myFile, err = os.OpenFile(dst, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("can't open this file")
		return ""
	}

	defer myFile.Close()
	//先读出内容
	var strSum string
	var read *bufio.Reader = bufio.NewReader(myFile)
	var str string
	for err != io.EOF {
		str, err = read.ReadString('\n')
		strSum += str
	}
	return strSum
}