package endpoints

import (
	"net/http"
	"ras/management/modems/sim"

	"github.com/gin-gonic/gin"
)

func SimInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		var request struct {
			Sim string `json:"sim"`
		}
		if err := c.ShouldBindBodyWithJSON(&request); err != nil || request.Sim == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "No sim specified",
			})
			return
		}

		info, err := sim.Get(request.Sim)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad sim specified",
			})
			return
		}

		c.JSON(http.StatusOK, info)
	}
}
