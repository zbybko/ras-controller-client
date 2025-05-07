package endpoints

import (
	"errors"
	"net/http"
	"ras/management/time"
	"ras/management/time/chrony"
	"slices"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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

func SetTimezoneHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			Timezone string `json:"timezone"`
		}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			log.Errorf("Error while binding timezone: %s", err)
			ctx.AbortWithStatus(http.StatusBadRequest)
		}
		err := time.SetTimeZone(request.Timezone)
		if err != nil {
			log.Errorf("Error while setting timezone: %s", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}

		ctx.JSON(http.StatusOK, gin.H{"timezone": request.Timezone})
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
			r.Servers = append(r.Servers, NtpServer{Address: ntp.Address(), Options: ntp.Options()})
		}
		ctx.JSON(http.StatusOK, r)
	}
}

type addNtpServerRequest struct {
	Address string   `json:"address"`
	Options []string `json:"options"`
}

func AddNtpServerHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var r addNtpServerRequest
		if err := ctx.ShouldBindJSON(&r); err != nil {
			log.Errorf("Error while binding request: %s", err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		server := chrony.NewNtpServer(r.Address)
		server.ChronyParameter.Options = r.Options

		// Adding default options
		defaultOptions := viper.GetStringSlice("ntp.default_options")
		for _, opt := range defaultOptions {
			if slices.Contains(server.ChronyParameter.Options, opt) {
				continue
			}
			server.ChronyParameter.Options = append(server.ChronyParameter.Options, opt)
		}

		err := time.AddNtpServer(server)
		if err != nil {
			log.Errorf("Error while adding NTP server: %s", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		ctx.Status(http.StatusOK)
	}
}

type removeNtpServerRequest struct {
	Address string `json:"address"`
}

func RemoveNtpServerHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var r removeNtpServerRequest
		if err := ctx.ShouldBindJSON(&r); err != nil {
			log.Errorf("Error while binding request: %s", err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		server := chrony.NewNtpServer(r.Address)
		err := time.RemoveNtpServer(*server)
		if err != nil {
			log.Errorf("Error while removing NTP server: %s", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		ctx.Status(http.StatusOK)
	}
}
