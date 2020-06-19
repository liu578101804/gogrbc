package sysinit

import (
	_ "github.com/liu578101804/gorbac/models"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func dbInit() {
	//如实是开发模式，则显示命令信息
	isDev := "dev" == beego.AppConfig.String("runmode")
	if isDev {
		orm.Debug = isDev
	}

	registerDatabase()

	_ = orm.RunSyncdb("default", false, isDev)
}

func registerDatabase() {

	alias := "default"

	dbName 	:= beego.AppConfig.String("db_database")
	dbUser 	:= beego.AppConfig.String("db_username")
	dbPwd 	:= beego.AppConfig.String("db_password")
	dbHost 	:= beego.AppConfig.String("db_host")
	dbPort 	:= beego.AppConfig.String("db_port")

	err := orm.RegisterDataBase(alias, "mysql", dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8", 30)

	if err != nil {
		beego.Error(err)
	}
}
