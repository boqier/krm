package pod

import (
	"github.com/boqier/krm/controllers/pod"
	"github.com/gin-gonic/gin"
)

func Create(r *gin.RouterGroup) {
	r.POST("/create", pod.Create)
	r.GET("/create", pod.Create)
}
func Update(r *gin.RouterGroup) {
	r.POST("/update", pod.Update)
	r.GET("/update", pod.Update)
}
func Delete(r *gin.RouterGroup) {
	r.POST("/delete", pod.Delete)
	r.GET("/delete", pod.Delete)
}
func Get(r *gin.RouterGroup) {
	r.POST("/get", pod.Get)
	r.GET("/get", pod.Get)
}
func List(r *gin.RouterGroup) {
	r.POST("/list", pod.List)
	r.GET("/list", pod.List)
}
func RegisterSubRouters(r *gin.RouterGroup) {
	//登录
	podGroup := r.Group("/pod")
	Create(podGroup)
	Update(podGroup)
	Delete(podGroup)
	Get(podGroup)
	List(podGroup)
}
