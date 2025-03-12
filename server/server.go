package server

import (
	"fmt"
	"net/http"
	"ras/management"
	"ras/management/time"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func setMode() {
	mode := viper.GetString("server.mode")
	if viper.GetBool("debug") {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
}

func New() *gin.Engine {
	setMode()
	srv := gin.Default()
	api := srv.Group("/api")

	api.GET("/os-info", func(ctx *gin.Context) {
		info, err := management.GetOSInfo()
		if err != nil {
			log.Errorf("Error while getting OS info: %s", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		ctx.JSON(http.StatusOK, info)
	})
	api.GET("/timezone", func(ctx *gin.Context) {
		zone, err := time.GetTimeZone()
		if err != nil {
			log.Errorf("Error while getting timezone: %s", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		ctx.JSON(http.StatusOK, gin.H{"timezone": zone})
	})

	return srv
}

func Address() string {
	return fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port"))
}
