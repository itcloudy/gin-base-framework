package middles

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func RequestLanguages() gin.HandlerFunc  {
	return func(c *gin.Context) {
		langHeader := c.Request.Header.Get("Accept-Language")
		fmt.Println(langHeader)

		c.Next()
	}
}