package utils

import (
	"github.com/astaxie/beego"
	"log"
	"os"
	"time"
)

func init() {
	//获取当天时间
	now := time.Now().Format("2006-01-02")
	file, err := os.OpenFile("/home/rzj/pub/git/mygo/src/d.rzj.me/log/"+now+".log", os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fd, err := os.Create("/home/rzj/pub/git/mygo/src/d.rzj.me/log/" + now + ".log")
		if err != nil {
			beego.Critical("openfile beepkg.log:", err)
			return
		}
		lg := log.New(fd, "", log.Ldate|log.Ltime)
		beego.SetLogger(lg)
	}
	lg := log.New(file, "", log.Ldate|log.Ltime)
	beego.SetLogger(lg)
}

func Info(v ...interface{}) {
	beego.Info(v)
}
