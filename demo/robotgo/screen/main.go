package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/go-vgo/robotgo"
)

// 可以做什么要考自己思考
func main() {
	x, y := robotgo.GetScreenSize()
	logs.Info("屏幕大小:", x, y)
	color := robotgo.GetPixelColor(x/2, y/2)
	logs.Info("屏幕:", x/2, y/2, " 颜色:", color)
	robotgo.CaptureScreen()
}
