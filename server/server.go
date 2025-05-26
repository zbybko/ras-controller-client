package server

import (
	"fmt"
	"net/http"
	"ras/config"
	"ras/management"
	"ras/management/nmcli"
	"ras/server/endpoints"
	"ras/server/endpoints/server_management"
	"ras/server/middleware"

	"github.com/charmbracelet/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func setMode() {
	config.SetupServerMode()
}

func New() *gin.Engine {
	log.Debugf("Initializing gin %s", gin.Version)
	setMode()
	srv := gin.New()

	srv.Use(middleware.Logger())
	srv.Use(cors.New(cors.Config{
		AllowOrigins:     viper.GetStringSlice("client.addresses"),
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	api := srv.Group("/api")
	if gin.IsDebugging() {
		api.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status": gin.IsDebugging(),
			})
		})
	}
	api.GET("/os-info", func(ctx *gin.Context) {
		info, err := management.GetOSInfo()
		if err != nil {
			log.Errorf("Error while getting OS info: %s", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		ctx.JSON(http.StatusOK, info)
	})
	api.GET("/netload", endpoints.NetloadHandler())
	api.GET("/device-info", endpoints.DeviceInfoHandler())
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
		band2 := wifi.Group("/2")
		{
			band2.POST("/enable", endpoints.EnableWiFi(nmcli.WirelessBand2GHz))
			band2.POST("/disable", endpoints.DisableWiFi(nmcli.WirelessBand2GHz))
			band2.GET("/status", endpoints.WiFiStatus(nmcli.WirelessBand2GHz))
			band2.POST("/ssid/hide", endpoints.SetSSIDHidden(nmcli.WirelessBand2GHz))
			band2.POST("/ssid/set", endpoints.SetSSID(nmcli.WirelessBand2GHz))
			band2.POST("/password/set", endpoints.SetPassword(nmcli.WirelessBand2GHz))
			band2.POST("/channel/set", endpoints.SetChannel(nmcli.WirelessBand2GHz))
		}
		band5 := wifi.Group("/5")
		{
			band5.POST("/enable", endpoints.EnableWiFi(nmcli.WirelessBand5GHz))
			band5.POST("/disable", endpoints.DisableWiFi(nmcli.WirelessBand5GHz))
			band5.GET("/status", endpoints.WiFiStatus(nmcli.WirelessBand5GHz))
			band5.POST("/ssid/hide", endpoints.SetSSIDHidden(nmcli.WirelessBand5GHz))
			band5.POST("/ssid/set", endpoints.SetSSID(nmcli.WirelessBand5GHz))
			band5.POST("/password/set", endpoints.SetPassword(nmcli.WirelessBand5GHz))
			band5.POST("/channel/set", endpoints.SetChannel(nmcli.WirelessBand5GHz))
		}

		wifi.GET("/connected-clients", endpoints.ConnectedClientsHandler())
	}
	auth := api.Group("/auth")
	{
		auth.POST("/signin", endpoints.Authorization())
		auth.POST("/change-password", endpoints.ChangePassword())
	}
	modems := api.Group("/modems")
	{
		modems.GET("/list", endpoints.ModemsListHandler())
		modems.POST("/disable/:modem", endpoints.DisableModemHandler())
		modems.POST("/enable/:modem", endpoints.EnableModemHandler())
		modems.GET("/signal/:modem", endpoints.GetModemSignalHandlers()...)
	}
	dhcp := api.Group("/dhcp")
	{
		dhcp.GET("/can-manage", endpoints.CanManageDhcpHandler())
		dhcp.GET("/status", endpoints.DhcpStatusHandler())
		dhcp.POST("/enable", endpoints.EnableDhcpHandler())
		dhcp.POST("/disable", endpoints.DisableDhcpHandler())
		dhcp.GET("/leases", endpoints.LeasesDhcpHandler())
		dhcp.GET("/ranges", endpoints.GetDhcpRangeHandler())
		dhcp.POST("/ranges", endpoints.SetDhcpRangeHandler())

		static := dhcp.Group("/static")
		{
			static.GET("/list", endpoints.GetStaticLeasesHandler())
			static.POST("/add", endpoints.AddStaticLeaseHandler())
			static.POST("/remove", endpoints.RemoveStaticLeaseHandler())
		}
	}
	api.GET("/journal/:journal", endpoints.JournalsHandler())
	api.POST("/sim/:sim", endpoints.SimInfoHandler())
	api.POST("/reboot", endpoints.RebootHandler())
	ssh := api.Group("/ssh")
	{
		ssh.GET("/status", endpoints.SshStatusHandler())
		ssh.POST("/enable", endpoints.EnableSshHandler())
		ssh.POST("/disable", endpoints.DisableSshHandler())
	}
	diag := api.Group("/diagnostics")
	{
		diag.GET("/default-ping-address", endpoints.DefaultPingAddressHandler())
		diag.POST("/ping/:address", endpoints.PingHandler())
		diag.POST("/nslookup/:address", endpoints.NslookupHandler())
		diag.POST("/traceroute/:address", endpoints.TracerouteHandler())
	}
	ethernet := api.Group("/ethernet")
	{
		ethernet.GET("/status", endpoints.GetEthernetPortsHandler())
	}
	ras := api.Group("/ras")
	{
		server := ras.Group("/server")
		server.POST("/mode/:mode", server_management.SetModeHandler())
		server.GET("/mode", server_management.GetModeHandler())
	}
	return srv
}

func Address() string {
	return fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port"))
}
