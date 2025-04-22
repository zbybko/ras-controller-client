package endpoints

import (
	"net/http"
	"ras/management/ethernet"

	"github.com/gin-gonic/gin"
)

func GetEthernetPortsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ports, err := ethernet.GetEthernetPorts()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, ports)
	}
}
