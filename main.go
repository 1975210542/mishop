package main

import (
	"encoding/gob"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	"mi.com/models"
	_ "mi.com/routers"
)

func init() {
	gob.Register(models.Manager{})
}

func main() {
	// 注册模板函数
	_ = beego.AddFuncMap("unixToDate", models.UnixToDate)
	_ = beego.AddFuncMap("formatImg", models.FormatImg)
	_ = beego.AddFuncMap("formatAttr", models.FormatAttr)
	_ = beego.AddFuncMap("cutStr", models.CutStr)
	_ = beego.AddFuncMap("mul", models.Mul)
	_ = beego.AddFuncMap("judge", models.Judge)

	// 配置session保存在redis里面
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"
	beego.Run()
	defer models.DB.Close()
}
