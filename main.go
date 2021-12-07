package main

import (
	_ "charroom/routers"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

const (
	APP_VER = "0.1.1.0227"
)

func main() {
	beego.Info(beego.BConfig.AppName, APP_VER)
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()
}
