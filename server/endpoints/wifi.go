package endpoints

import (
	"net/http"
	"ras/management/wifi"
	"ras/management/wifi/hostapd"

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
		status, err := wifi.Status()
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, status)
	}
}

// Скрыть/показать SSID
func SetSSIDHidden(c *gin.Context) {
	var req struct {
		Hidden bool `json:"hidden"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := hostapd.SetSSIDHidden(req.Hidden); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	returnWiFiStatus(c)
}

// Изменить имя сети (SSID)
func SetSSID(c *gin.Context) {
	var req struct {
		SSID string `json:"ssid"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := hostapd.SetSSID(req.SSID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	returnWiFiStatus(c)
}

// Изменить пароль сети
func SetPassword(c *gin.Context) {
	var req struct {
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := hostapd.SetPassword(req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	returnWiFiStatus(c)
}

// Изменить тип шифрования (WPA2/WPA3)
func SetSecurityType(c *gin.Context) {
	var req struct {
		WPA3 bool `json:"wpa3"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := hostapd.SetSecurityType(req.WPA3); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	returnWiFiStatus(c)
}

// Установить Wi-Fi канал
func SetChannel(c *gin.Context) {
	var req struct {
		Channel int `json:"channel"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := hostapd.SetChannel(req.Channel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	returnWiFiStatus(c)
}

// Вспомогательная функция для возврата статуса Wi-Fi
func returnWiFiStatus(c *gin.Context) {
	status, err := wifi.Status()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, status)
}
