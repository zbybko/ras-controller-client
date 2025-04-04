package endpoints

import (
	"net/http"
	"ras/management/dhcp"

	"github.com/gin-gonic/gin"
)

func DhcpStatusHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		status := dhcp.Status()
		ctx.JSON(http.StatusOK, status)
	}
}
func EnableDhcpHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := dhcp.Enable()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"enabled": true})
	}
}
func DisableDhcpHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := dhcp.Disable()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"enabled": false})
	}
}
func LeasesDhcpHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		leases, err := dhcp.GetLeases()
		if err == nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		ctx.JSON(http.StatusOK, leases)
	}
}
