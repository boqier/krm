package main

import (
	"github.com/boqier/krm/config"
	_ "github.com/boqier/krm/controllers/initcontroller"
	"github.com/boqier/krm/middlerwares"
	"github.com/boqier/krm/routers"
	"github.com/boqier/krm/utils/logs"
	gin "github.com/gin-gonic/gin"
)

func main() {
	//1.加载配置
	r := gin.Default()
	r.Use(middlerwares.JWTAuth)
	routers.RegisterRouters(r)
	logs.Info(map[string]interface{}{"port": config.Port}, "krm server start at :%d")
	r.Run(":" + config.Port)
}
