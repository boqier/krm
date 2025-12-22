package namespace

import (
	"github.com/boqier/krm/controllers/namespace"
	"github.com/gin-gonic/gin"
)

func Add(r *gin.RouterGroup) {
	r.POST("/add", namespace.Add)
}
func Update(r *gin.RouterGroup) {
	r.POST("/update", namespace.Update)
}
func Delete(r *gin.RouterGroup) {
	r.POST("/delete", namespace.Delete)
}
func Get(r *gin.RouterGroup) {
	r.POST("/get", namespace.Get)
}
func List(r *gin.RouterGroup) {
	r.POST("/list", namespace.List)
}
func RegisterSubRouters(r *gin.RouterGroup) {
	//登录
	namespaceGroup := r.Group("/namespace")
	Add(namespaceGroup)
	Update(namespaceGroup)
	Delete(namespaceGroup)
	Get(namespaceGroup)
	List(namespaceGroup)
}
