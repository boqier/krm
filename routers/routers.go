package routers

import (
	"github.com/boqier/gin-scaffold/routers/auth"
	"github.com/gin-gonic/gin"
)

// 注册路由
func RegisterRouters(r *gin.Engine) {
	//登录
	apiGroup := r.Group("/api")
	auth.RegisterAuthRouters(apiGroup)

}
