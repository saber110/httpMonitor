package routers

import (
	"github.com/astaxie/beego"
	"github.com/saber110/httpMonitor/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "*:Index")
	beego.AutoRouter(&controllers.IndexController{})
	beego.AutoRouter(&controllers.GroupController{})
	beego.AutoRouter(&controllers.JobController{})
	beego.AutoRouter(&controllers.ConfigController{})
	beego.AutoRouter(&controllers.LoginController{})
}
