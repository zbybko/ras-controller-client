package endpoints

import (
	"net/http"
	"ras/management/modems/sim"

	"github.com/gin-gonic/gin"
)

func SimInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		simName, found := c.Params.Get("sim")

		if !found || simName == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid sim specified",
			})
			return
		}

		info, err := sim.Get(simName)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad sim specified",
			})
			return
		}

		c.JSON(http.StatusOK, info)
	}
}
