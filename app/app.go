package main

import (
	_ "github.com/edwinakatosh/beego-security-headers/app/routers"

	"github.com/beego/beego/v2/server/web"
	securityHeaders "github.com/edwinakatosh/beego-security-headers"
)

func main() {
	securityHeaders.Init()
	web.Run()
}
