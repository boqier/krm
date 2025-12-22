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

func Delete(ctx *gin.Context) {
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
	err = clientSet.CoreV1().Namespaces().Delete(context.TODO(), namespace.Name, metav1.DeleteOptions{})
	if err != nil {
		returnData := config.NewReturnDate()
		returnData.Status = 400
		logs.Error(map[string]interface{}{"namespace": namespace.Name}, err.Error())
		returnData.Message = "删除namespace失败"
		ctx.JSON(200, returnData)
		return
	}
	returnData := config.NewReturnDate()
	returnData.Status = 200
	returnData.Message = "删除命名空间成功"
	ctx.JSON(200, returnData)
}
