package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xlgwr/demoApigo/controllers"
	_ "github.com/xlgwr/demoApigo/routers"
)

func init() {
	runmode := beego.AppConfig.String("runmode")
	mysqlurls := beego.AppConfig.String(runmode + "::mysqlurls")
	fmt.Println(mysqlurls)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	maxIdle := beego.AppConfig.DefaultInt("maxIdle", 30)
	maxConn := beego.AppConfig.DefaultInt("maxConn", 30)
	orm.RegisterDataBase("default", "mysql", mysqlurls, maxIdle, maxConn)

	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
