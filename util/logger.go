package util

import (
	"bingWallpaper/constant"
	"io"
	"log"
	"os"
)

var(
	Info 	*log.Logger
	Error 	*log.Logger
)

func init()  {
	file, _ := os.OpenFile(constant.Log, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	//自定义日志格式
	Info = log.New(io.MultiWriter(file, os.Stderr), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}