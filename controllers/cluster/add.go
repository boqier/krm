package cluster

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Add(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "add cluster",
	})
}
