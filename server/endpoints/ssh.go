package endpoints

import (
	"net/http"
	"ras/management/ssh"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func sshStatus(ctx *gin.Context) {
	status := ssh.Status()
	ctx.JSON(http.StatusOK, status)
}

func DisableSshHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := ssh.Disable()
		if err != nil {
			log.Errorf("Failed disable ssh: %s", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		sshStatus(ctx)
	}
}
func EnableSshHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := ssh.Enable()
		if err != nil {
			log.Errorf("Failed enable ssh: %s", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		sshStatus(ctx)
	}
}

func SshStatusHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sshStatus(ctx)
	}
}
