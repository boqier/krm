package pod

import (
	"context"

	"github.com/boqier/krm/config"
	"github.com/boqier/krm/controllers"
	"github.com/boqier/krm/utils/logs"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Delete(ctx *gin.Context) {
	//获取客户端
	clientSet, basicInfo, err := controllers.GetClientSet(ctx, nil)
	if err != nil {
		returnData := config.NewReturnDate()
		returnData.Status = 400
		returnData.Message = err.Error()
		ctx.JSON(200, returnData)
		return
	}
	logs.Warning(map[string]interface{}{"pod列表": basicInfo.DeleteList}, "删除pod")
	for _, podName := range basicInfo.DeleteList {
		_ = clientSet.CoreV1().Pods(basicInfo.Namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})

	}
	returnData := config.NewReturnDate()
	returnData.Status = 200
	returnData.Message = "删除pod成功"
	ctx.JSON(200, returnData)
}
