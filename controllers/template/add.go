package temp

import (
	"github.com/boqier/krm/config"
	"github.com/gin-gonic/gin"
)

func Add(ctx *gin.Context) {
	returnData := config.NewReturnDate()
	returnData.Status = 200
	returnData.Message = "添加模板成功"
	ctx.JSON(200, returnData)

}
