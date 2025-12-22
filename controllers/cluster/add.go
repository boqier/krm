package cluster

import (
	"github.com/gin-gonic/gin"
)

func Add(ctx *gin.Context) {
	addOrUpdate(ctx, "create")

}
