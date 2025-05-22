package endpoints

import (
	"fmt"
	"net/http"
	"ras/management/wifi"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

// Включить Wi-Fi
func EnableWiFi(c *gin.Context) {
	if err := wifi.Enable(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	returnWiFiStatus(c)
}

// Выключить Wi-Fi
func DisableWiFi(c *gin.Context) {
	if err := wifi.Disable(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	returnWiFiStatus(c)
}

// Получить статус Wi-Fi
func WiFiStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		returnWiFiStatus(ctx)
	}
}

// Скрыть/показать SSID
func SetSSIDHidden(c *gin.Context) {
	var req struct {
		Hidden bool `json:"hidden" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := wifi.SetHidden(req.Hidden); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	returnWiFiStatus(c)
}

// Изменить имя сети (SSID)
func SetSSID(c *gin.Context) {
	var req struct {
		SSID string `json:"ssid" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := wifi.SetSSID(req.SSID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	returnWiFiStatus(c)
}

// Изменить пароль сети
func SetPassword(c *gin.Context) {
	var req struct {
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request: %s", err)})
		return
	}

	log.Debugf("[ENDPOINT] Setting password to '%s'", req.Password)

	if err := wifi.SetPassword(req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	returnWiFiStatus(c)
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
func SetChannel(c *gin.Context) {
	var req struct {
		Channel int `json:"channel"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := wifi.SetChannel(req.Channel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	returnWiFiStatus(c)
}

// Вспомогательная функция для возврата статуса Wi-Fi
func returnWiFiStatus(c *gin.Context) {
	status, err := wifi.Status()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, status)
}
