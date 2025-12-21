package cluster

import (
	"github.com/boqier/krm/controllers/cluster"
	"github.com/gin-gonic/gin"
)

func Add(r *gin.RouterGroup) {
	r.POST("/add", cluster.Add)
}
func Update(r *gin.RouterGroup) {
	r.POST("/update", cluster.Update)
}
func Delete(r *gin.RouterGroup) {
	r.POST("/delete", cluster.Delete)
}
func Get(r *gin.RouterGroup) {
	r.POST("/get", cluster.Get)
}
func List(r *gin.RouterGroup) {
	r.POST("/list", cluster.List)
}
func RegisterSubRouters(r *gin.RouterGroup) {
	//登录
	clusterGroup := r.Group("/cluster")
	Add(clusterGroup)
	Update(clusterGroup)
	Delete(clusterGroup)
	Get(clusterGroup)
	List(clusterGroup)
}
