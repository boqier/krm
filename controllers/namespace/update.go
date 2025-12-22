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

func Update(ctx *gin.Context) {
	//获取客户端
	var ns corev1.Namespace

	clientSet, _, err := controllers.GetClientSet(ctx, &ns)
	if err != nil {
		returnData := config.NewReturnDate()
		returnData.Status = 400
		returnData.Message = err.Error()
		ctx.JSON(200, returnData)
		return
	}

	_, err = clientSet.CoreV1().Namespaces().Update(context.TODO(), &ns, metav1.UpdateOptions{})
	if err != nil {
		returnData := config.NewReturnDate()
		returnData.Status = 400
		logs.Error(map[string]interface{}{"namespace": ns.Name}, err.Error())
		returnData.Message = "更新命名空间失败，更新命名空间失败"
		ctx.JSON(200, returnData)
		return
	}
	returnData := config.NewReturnDate()
	returnData.Status = 200
	returnData.Message = "更新命名空间成功"
	ctx.JSON(200, returnData)
}
