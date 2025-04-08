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
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		ctx.JSON(http.StatusOK, leases)
	}
}
func SetDhcpRangeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req struct {
			Subnet            string `json:"subnet"`
			Netmask           string `json:"netmask"`
			StartIP           string `json:"start_ip"`
			EndIP             string `json:"end_ip"`
			OptionsRouters    string `json:"options_routers"`
			OptionsBroadcasts string `json:"options_broadcasts"`
		}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request. Start IP and End IP are required."})
			return
		}

		err := dhcp.SetDhcpRange(req.Subnet, req.Netmask, req.StartIP, req.EndIP, req.OptionsRouters, req.OptionsBroadcasts)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = dhcp.RestartDhcp()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to restart DHCP service"})
			return
		}

		currentRange, err := dhcp.GetDhcpRange()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get current DHCP range"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"current_range": currentRange,
		})
	}
}
func GetDhcpRangeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		currentRange, err := dhcp.GetDhcpRange()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get current DHCP range"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"current_range": currentRange,
		})
	}
}
