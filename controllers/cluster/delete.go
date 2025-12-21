package cluster

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete cluster",
	})
}
