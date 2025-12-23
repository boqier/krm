package routers

import (
	"github.com/boqier/krm/routers/auth"
	"github.com/boqier/krm/routers/cluster"
	"github.com/boqier/krm/routers/namespace"
	"github.com/boqier/krm/routers/pod"
	"github.com/gin-gonic/gin"
)

// 注册路由
func RegisterRouters(r *gin.Engine) {
	//登录
	apiGroup := r.Group("/api")
	auth.RegisterSubRouters(apiGroup)
	cluster.RegisterSubRouters(apiGroup)
	namespace.RegisterSubRouters(apiGroup)
	pod.RegisterSubRouters(apiGroup)
}
