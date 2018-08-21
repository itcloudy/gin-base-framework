package middles

import (
	"github.com/gin-gonic/gin"
	"github.com/hexiaoyun128/gin-base-framework/common"
)

func VisitHistory() gin.HandlerFunc {
	return func(c *gin.Context) {
		addr:=c.Request.RemoteAddr
		common.Logger.Info("visit address"+addr)
	}
}
