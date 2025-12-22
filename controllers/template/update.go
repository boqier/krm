package namespace

import (
	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context) {
	addOrUpdate(ctx, "update")
}
