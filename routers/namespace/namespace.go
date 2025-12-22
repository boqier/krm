package namespace

import (
	"github.com/boqier/krm/controllers/namespace"
	"github.com/gin-gonic/gin"
)

func Create(r *gin.RouterGroup) {
	r.POST("/create", namespace.Create)
	r.GET("/create", namespace.Create)
}
func Update(r *gin.RouterGroup) {
	r.POST("/update", namespace.Update)
	r.GET("/update", namespace.Update)
}
func Delete(r *gin.RouterGroup) {
	r.POST("/delete", namespace.Delete)
	r.GET("/delete", namespace.Delete)
}
func Get(r *gin.RouterGroup) {
	r.POST("/get", namespace.Get)
	r.GET("/get", namespace.Get)
}
func List(r *gin.RouterGroup) {
	r.POST("/list", namespace.List)
	r.GET("/list", namespace.List)
}
func RegisterSubRouters(r *gin.RouterGroup) {
	//登录
	namespaceGroup := r.Group("/namespace")
	Create(namespaceGroup)
	Update(namespaceGroup)
	Delete(namespaceGroup)
	Get(namespaceGroup)
	List(namespaceGroup)
}
