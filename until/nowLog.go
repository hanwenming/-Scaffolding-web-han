package until

import (
	blog "github.com/higker/logker"
)

//日志初始化
func Log_init() (blog.Logger, error) {
	// 在初始化的时候必须自己先创建好存储日志的文件夹!!
	dir := "../log"
	// New file logger
	// File Max size : 这里是单个的日志文件大小 你可以自定义也可以使用内置的常量
	// blog.GB1  	= 1GB
	// blog.MB10  	= 10MB
	// blog.MB100	= 100MB
	format := "{level} - 时间 {time}  - 位置 {position} - 消息 {message}"
	//创建日志缓冲区
	var Qs1w int64 = 1000
	task := blog.InitAsync(Qs1w) //This version was modified from v 1.1.8
	flog, err := blog.NewFlog(blog.DEBUG, true, blog.Shanghai, dir, "log", 10*1024, 0777, format, task)
	return flog, err
}
