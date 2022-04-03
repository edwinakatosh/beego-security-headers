package routers

import (
	"github.com/edwinakatosh/beego-security-headers/app/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.MainController{}, "get:Home")
}
