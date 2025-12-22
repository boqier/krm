package namespace

import (
	"context"

	"github.com/boqier/krm/config"
	"github.com/boqier/krm/controllers"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func List(ctx *gin.Context) {
	//获取客户端
	clientSet, _, err := controllers.GetClientSet(ctx, nil)
	if err != nil {
		returnData := config.NewReturnDate()
		returnData.Status = 400
		returnData.Message = err.Error()
		ctx.JSON(200, returnData)
		return
	}

	NamespaceList, err := clientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		returnData := config.NewReturnDate()
		returnData.Status = 400
		returnData.Message = err.Error()
		ctx.JSON(200, returnData)
		return
	}
	returnData := config.NewReturnDate()
	returnData.Status = 200
	returnData.Message = "获取命名空间成功"
	returnData.Data = map[string]interface{}{
		"namespaceList": NamespaceList.Items,
	}
	ctx.JSON(200, returnData)
}
