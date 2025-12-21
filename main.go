package main

import (
	"github.com/boqier/gin-scaffold/config"
	"github.com/boqier/gin-scaffold/middlerwares"
	"github.com/boqier/gin-scaffold/routers"
	"github.com/boqier/gin-scaffold/utils/logs"
	"github.com/gin-gonic/gin"
)

func main() {
	logs.Info(map[string]interface{}{"module": "main"}, "正在启动服务...")
	r := gin.Default()
	//注册中间件
	r.Use(middlerwares.JWTAuth)
	routers.RegisterRouters(r)
	r.Run("0.0.0.0:" + config.Port)
}
