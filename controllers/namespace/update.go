package namespace

import (
	"github.com/boqier/krm/config"
	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context) {
	returnData := config.NewReturnDate()
	returnData.Status = 200
	returnData.Message = "更新命名空间成功"
	ctx.JSON(200, returnData)
}
