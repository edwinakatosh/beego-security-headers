package main

import (
	"github.com/beego/beego/v2/server/web"
	securityHeaders "github.com/edwinakatosh/beego-security-headers"
)

func main() {
	securityHeaders.Init()
	web.Run()
}
