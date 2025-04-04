package server

import (
	"fmt"
	"net/http"
	"ras/management"
	"ras/server/endpoints"
	"ras/server/middleware"

	"github.com/charmbracelet/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func setMode() {
	mode := viper.GetString("server.mode")
	if viper.GetBool("debug") {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
	log.Debugf("Current gin mode: '%s'", gin.Mode())
}

func New() *gin.Engine {
	log.Debugf("Initializing gin %s", gin.Version)
	setMode()
	srv := gin.New()

	srv.Use(middleware.Logger())
	srv.Use(cors.New(cors.Config{
		AllowOrigins:     []string{viper.GetString("client.address")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

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
	firewall := api.Group("/firewall")
	{
		firewall.POST("/enable", endpoints.EnableFirewall)
		firewall.POST("/disable", endpoints.DisableFirewall)
		firewall.GET("/status", endpoints.FirewallStatus())
		firewall.GET("/can-manage", endpoints.CanManageFirewallHandler())
	}
	wifi := api.Group("/wifi")
	{
		wifi.POST("/enable", endpoints.EnableWiFi)
		wifi.POST("/disable", endpoints.DisableWiFi)
		wifi.GET("/status", endpoints.WiFiStatus())
		wifi.POST("/ssid/hide", endpoints.SetSSIDHidden)
		wifi.POST("/ssid/set", endpoints.SetSSID)
		wifi.POST("/password/set", endpoints.SetPassword)
		wifi.POST("/security/set", endpoints.SetSecurityType)
		wifi.POST("/channel/set", endpoints.SetChannel)
	}
	api.POST("/auth/", endpoints.Authorization())
	modems := api.Group("/modems")
	{
		modems.GET("/list", endpoints.ModemsListHandler())
		modems.POST("/disable/:modem", endpoints.DisableModemHandler())
		modems.POST("/enable/:modem", endpoints.EnableModemHandler())
		modems.GET("/signal/:modem", endpoints.GetModemSignalHandlers()...)
	}
	dhcp := api.Group("/dhcp")
	{
		dhcp.GET("/status", endpoints.DhcpStatusHandler())
		dhcp.POST("/enable", endpoints.EnableDhcpHandler())
		dhcp.POST("/disable", endpoints.DisableDhcpHandler())
		dhcp.GET("/leases", endpoints.LeasesDhcpHandler())
	}
	api.GET("/journal/:journal", endpoints.JournalsHandler())
	api.POST("/sim/:sim", endpoints.SimInfoHandler())
	api.POST("/reboot", endpoints.RebootHandler())
	return srv
}

func Address() string {
	return fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port"))
}
