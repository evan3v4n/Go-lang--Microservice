// Path: project_root/api/handlers/error_handler.go

package handlers

import (
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			c.JSON(c.Writer.Status(), gin.H{
				"errors": c.Errors.Errors(),
			})
		}
	}
}
