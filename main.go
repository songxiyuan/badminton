package main

import (
	"badminton/third"
	"badminton/util"
	"log"
	"os"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	logFile, err := os.OpenFile("./log/temp.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开日志文件异常")
	}
	log.SetOutput(logFile)
}

func main() {
	monitor()
	time.Sleep(time.Minute)
}

func monitor() {
	var zyl third.ZYL
	res, err := zyl.Monitor([]string{
		"",
	})
	if err != nil {
		util.SendEmail("error")
		log.Println(err.Error())
		return
	}
	if res {
		util.SendEmail("有!")
		log.Println("有")
		return
	}
}
