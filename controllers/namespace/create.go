package namespace

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
	clientSet, basicInfo, err := controllers.GetClientSet(ctx)
	if err != nil {
		returnData := config.NewReturnDate()
		returnData.Status = 400
		returnData.Message = err.Error()
		ctx.JSON(200, returnData)
		return
	}
	var namespace corev1.Namespace
	namespace.Name = basicInfo.Name
	_, err = clientSet.CoreV1().Namespaces().Create(context.TODO(), &namespace, metav1.CreateOptions{})
	if err != nil {
		returnData := config.NewReturnDate()
		returnData.Status = 400
		logs.Error(map[string]interface{}{"namespace": namespace.Name}, err.Error())
		returnData.Message = "创建命名空间失败，创建命名空间失败"
		ctx.JSON(200, returnData)
		return
	}
	returnData := config.NewReturnDate()
	returnData.Status = 200
	returnData.Message = "创建命名空间成功"
	ctx.JSON(200, returnData)
}
