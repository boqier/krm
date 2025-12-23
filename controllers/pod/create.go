package pod

import (
	"context"

	"github.com/boqier/krm/config"
	"github.com/boqier/krm/controllers"
	"github.com/boqier/krm/utils/logs"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Create(ctx *gin.Context) {

	//获取客户端
	var pod corev1.Pod
	clientSet, basicInfo, err := controllers.GetClientSet(ctx, &pod)
	if err != nil {
		returnData := config.NewReturnDate()
		returnData.Status = 400
		returnData.Message = err.Error()
		ctx.JSON(200, returnData)
		return
	}
	var namespace corev1.Namespace
	namespace.Name = basicInfo.Namespace
	_, err = clientSet.CoreV1().Pods(basicInfo.Namespace).Create(context.TODO(), &pod, metav1.CreateOptions{})
	if err != nil {
		returnData := config.NewReturnDate()
		returnData.Status = 400
		logs.Error(map[string]interface{}{"pod": pod.Name}, err.Error())
		returnData.Message = "创建pod失败，创建pod失败"
		ctx.JSON(200, returnData)
		return
	}
	returnData := config.NewReturnDate()
	returnData.Status = 200
	returnData.Message = "创建pod成功"
	ctx.JSON(200, returnData)
}
