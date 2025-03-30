package endpoints

import (
	"net/http"
	"ras/management/power"

	"github.com/gin-gonic/gin"
)

func RebootHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		power.Reboot()
		c.AbortWithStatus(http.StatusOK)
	}
}
