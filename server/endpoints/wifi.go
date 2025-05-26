package endpoints

import (
	"fmt"
	"net/http"
	"ras/management/iw"
	"ras/management/nmcli"
	"ras/management/wifi"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

// Включить Wi-Fi
func EnableWiFi(band nmcli.WirelessBand) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := wifi.Enable(band); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		returnWiFiStatus(c, band)
	}
}

// Выключить Wi-Fi
func DisableWiFi(band nmcli.WirelessBand) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := wifi.Disable(band); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		returnWiFiStatus(c, band)
	}
}

// Получить статус Wi-Fi
func WiFiStatus(band nmcli.WirelessBand) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		returnWiFiStatus(ctx, band)
	}
}

// Скрыть/показать SSID
func SetSSIDHidden(band nmcli.WirelessBand) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Hidden bool `json:"hidden" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		if err := wifi.SetHidden(band, req.Hidden); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		returnWiFiStatus(c, band)
	}
}

// Изменить имя сети (SSID)
func SetSSID(band nmcli.WirelessBand) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			SSID string `json:"ssid" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		if err := wifi.SetSSID(band, req.SSID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		returnWiFiStatus(c, band)
	}
}

// Изменить пароль сети
func SetPassword(band nmcli.WirelessBand) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request: %s", err)})
			return
		}

		log.Debugf("[ENDPOINT] Setting password to '%s'", req.Password)

		if err := wifi.SetPassword(band, req.Password); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		returnWiFiStatus(c, band)
	}
}

// Изменить тип шифрования (WPA2/WPA3)
// func SetSecurityType(c *gin.Context) {
// 	var req struct {
// 		WPA3 bool `json:"wpa3"`
// 	}

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
// 		return
// 	}

// 	var securityType hostapd.SecurityType

// 	if req.WPA3 {
// 		securityType = hostapd.SecurityTypeWPA3
// 	} else {
// 		securityType = hostapd.SecurityTypeWPA2
// 	}

// 	if err := hostapd.SetSecurityType(securityType); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	returnWiFiStatus(c)
// }

// Установить Wi-Fi канал
func SetChannel(band nmcli.WirelessBand) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Channel int `json:"channel"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		if err := wifi.SetChannel(band, req.Channel); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		returnWiFiStatus(c, band)
	}
}

func ConnectedClientsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clients, err := iw.GetConnectedDevices()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"success": false})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"clients": clients,
		})
	}
}

// Вспомогательная функция для возврата статуса Wi-Fi
func returnWiFiStatus(c *gin.Context, band nmcli.WirelessBand) {
	status, err := wifi.Status(band)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, status)
}
