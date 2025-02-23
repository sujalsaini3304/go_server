package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Author":  "Sujal Kumar Saini",
			"Message": "Go Server Started",
			"Status":  "O.K.",
		})
	})
}
