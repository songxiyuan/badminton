package main

import (
	"badminton/third"
	"badminton/util"
	"log"
	"math/rand"
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
	for {
		monitor()
		second := 50 + rand.Intn(20)
		time.Sleep(time.Duration(second) * time.Second)
	}
}

func monitor() {
	var zyl third.ZYL
	hit, url := zyl.Monitor([]string{
		"2022_10_05",
	})
	if hit {
		msg := "有! " + url
		util.SendEmail(msg)
		log.Println(msg)
		return
	}
	log.Println("无")
}
