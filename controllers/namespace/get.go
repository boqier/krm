package namespace

import (
	"github.com/boqier/krm/config"
	"github.com/boqier/krm/utils/logs"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Get(ctx *gin.Context) {
	logs.Info(map[string]interface{}{"集群ID": ctx.Param("id")}, "获取集群配置")
	clusterID := ctx.Query("clusterid")
	returnData := config.NewReturnDate()
	//获取集群配置
	secret, err := config.ClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Get(ctx, clusterID, metav1.GetOptions{})
	if err != nil {
		returnData.Status = 400
		returnData.Message = err.Error()
		ctx.JSON(400, returnData)
		logs.Error(map[string]interface{}{"集群ID": clusterID, "msg": err.Error()}, "获取集群配置失败")
		return
	} else {
		returnData.Status = 200
		returnData.Message = "success"
		returnData.Data = make(map[string]interface{})
		ClusterConfigMap := secret.Annotations
		ClusterConfigMap["kubeconfig"] = string(secret.Data["kubeconfig"])
		returnData.Data["item"] = ClusterConfigMap
		ctx.JSON(200, returnData)
	}

}
