package temp

import (
	"github.com/boqier/krm/config"
	"github.com/boqier/krm/utils/logs"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Delete(ctx *gin.Context) {
	logs.Debug(nil, "删除集群")
	clusterID := ctx.Query("clusterid")
	if clusterID == "" {
		logs.Error(nil, "删除集群失败,集群ID不能为空")
		returnData := config.NewReturnDate()
		returnData.Status = 400
		returnData.Message = "删除集群失败,集群ID不能为空"
		ctx.JSON(200, returnData)
		return
	}
	logs.Info(map[string]interface{}{"集群ID": clusterID}, "开始删除集群")
	err := config.ClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Delete(ctx, clusterID, metav1.DeleteOptions{})
	if err != nil {
		logs.Error(map[string]interface{}{"集群ID": clusterID}, "删除集群失败")
		returnData := config.NewReturnDate()
		returnData.Status = 400
		returnData.Message = "删除集群失败，请确认集群是否存在"
		ctx.JSON(200, returnData)
		return
	}
	returnData := config.NewReturnDate()
	returnData.Status = 200
	returnData.Message = "删除集群成功"
	ctx.JSON(200, returnData)
	logs.Info(map[string]interface{}{"集群ID": clusterID}, "删除集群成功")
}
