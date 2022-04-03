package controllers

import "github.com/beego/beego/v2/server/web"

// MainController:
// The controller must implement ControllerInterface
// Usually we extends web.Controller
type MainController struct {
	web.Controller
}

// address: http://localhost:8080 GET
func (ctrl *MainController) Home() {
	ctrl.Data["name"] = "Get()"
	ctrl.TplName = "home.tpl"
	// don't forget this
	_ = ctrl.Render()
}
