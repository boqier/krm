package pod

import (
	"context"

	"github.com/boqier/krm/config"
	"github.com/boqier/krm/controllers"
	"github.com/boqier/krm/utils/logs"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Get(ctx *gin.Context) {
	//获取客户端
	clientSet, basicInfo, err := controllers.GetClientSet(ctx, nil)
	if err != nil {
		returnData := config.NewReturnDate()
		returnData.Status = 400
		returnData.Message = err.Error()
		ctx.JSON(200, returnData)
		return
	}

	item, err := clientSet.CoreV1().Pods(basicInfo.Namespace).Get(context.TODO(), basicInfo.Name, metav1.GetOptions{})
	if err != nil {
		returnData := config.NewReturnDate()
		returnData.Status = 400
		logs.Error(map[string]interface{}{"pod": basicInfo.Name}, err.Error())
		returnData.Message = "获取pod失败，获取pod失败"
		ctx.JSON(200, returnData)
		return
	}
	returnData := config.NewReturnDate()
	returnData.Status = 200
	returnData.Message = "获取Pod成功"
	returnData.Data = map[string]interface{}{
		"item": item,
	}
	ctx.JSON(200, returnData)

}
