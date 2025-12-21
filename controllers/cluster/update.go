package cluster

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "update cluster",
	})
}
