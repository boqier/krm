package auth

import (
	"github.com/boqier/krm/controllers/auth"
	"github.com/gin-gonic/gin"
)

func login(authGroup *gin.RouterGroup) {
	authGroup.POST("/login", auth.Login)
	authGroup.GET("/register", auth.Login)
}
func logout(authGroup *gin.RouterGroup) {
	authGroup.POST("/logout", auth.Logout)
}

// 注册路由
func RegisterSubRouters(r *gin.RouterGroup) {
	//登录
	authGroup := r.Group("/auth")
	login(authGroup)
	logout(authGroup)
}
