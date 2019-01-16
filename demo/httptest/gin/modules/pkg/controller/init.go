package controller

import "github.com/gin-gonic/gin"

var r *gin.Engine

func init() {
	r = gin.Default()

}

// Start 启动WEB服务
func Start() {
	gin.SetMode(gin.ReleaseMode)
	r.Run()
}
