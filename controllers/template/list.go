package temp

import (
	"github.com/boqier/krm/config"
	"github.com/boqier/krm/utils/logs"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func List(ctx *gin.Context) {
	logs.Debug(nil, "获取集群列表")
	ListOptions := metav1.ListOptions{
		LabelSelector: "kubeasy.com/cluster.metadata=true",
	}
	secrets, err := config.ClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).List(ctx, ListOptions)
	if err != nil {
		logs.Error(nil, "获取集群列表失败")
		returnData := config.NewReturnDate()
		returnData.Status = 500
		returnData.Message = "获取集群列表失败"
		ctx.JSON(500, returnData)
		return
	}
	ClusterList := []map[string]string{}
	for _, v := range secrets.Items {
		anno := v.Annotations
		ClusterList = append(ClusterList, anno)
	}

	returnData := config.NewReturnDate()
	returnData.Status = 200
	returnData.Data["items"] = ClusterList
	returnData.Message = "获取集群列表成功"
	ctx.JSON(200, returnData)
}
