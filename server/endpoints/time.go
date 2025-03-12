package endpoints

import (
	"errors"
	"net/http"
	"ras/management/time"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func TimezoneHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		zone, err := time.GetTimeZone()
		if err != nil {
			log.Errorf("Error while getting timezone: %s", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		ctx.JSON(http.StatusOK, gin.H{"timezone": zone})
	}
}

type NtpServer struct {
	Address string   `json:"address"`
	Options []string `json:"options"`
}
type NtpInfoResponse struct {
	Active  bool        `json:"active"`
	Servers []NtpServer `json:"servers"`
}

func NtpInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		servers, err := time.GetNtpServers()
		active, err2 := time.IsNtpActive()
		if err = errors.Join(err, err2); err != nil {
			log.Errorf("Error while fetching NTP info: %s", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		r := NtpInfoResponse{
			Active: active,
		}
		for _, ntp := range servers {
			r.Servers = append(r.Servers, NtpServer{Address: ntp.Address(), Options: ntp.Options})
		}
		ctx.JSON(http.StatusOK, r)
	}
}
