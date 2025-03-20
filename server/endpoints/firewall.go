package endpoints

import (
	"net/http"
	"ras/management/firewall"

	"github.com/gin-gonic/gin"
)

func EnableFirewall(c *gin.Context) {
	if err := firewall.Enable(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	returnStatus(c)
}

func DisableFirewall(c *gin.Context) {
	if err := firewall.Disable(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	returnStatus(c)
}

func FirewallStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		status, err := firewall.Status()
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, status)
	}
}

func returnStatus(c *gin.Context) {
	status, err := firewall.Status()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, status)
}

func CanManageFirewallHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"canManage": firewall.ServiceExists()})
	}
}
