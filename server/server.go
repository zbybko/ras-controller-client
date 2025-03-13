package server

import (
	"fmt"
	"net/http"
	"ras/management"
	"ras/server/endpoints"
	"ras/server/middleware"

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
	srv := gin.New()
	srv.Use(middleware.Logger())
	api := srv.Group("/api")

	api.GET("/os-info", func(ctx *gin.Context) {
		info, err := management.GetOSInfo()
		if err != nil {
			log.Errorf("Error while getting OS info: %s", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		ctx.JSON(http.StatusOK, info)
	})
	api.GET("/timezone", endpoints.TimezoneHandler())
	api.POST("/timezone/set", endpoints.SetTimezoneHandler())
	api.GET("/ntp", endpoints.NtpInfo())
	api.POST("/ntp/add", endpoints.AddNtpServerHandler())
	api.DELETE("/ntp/remove", endpoints.RemoveNtpServerHandler())
	return srv
}

func Address() string {
	return fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port"))
}
