package endpoints

import (
	"net/http"
	"ras/management/modems"
	"ras/management/modems/sim"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

type ModemResponse struct {
	Operator string
	APN      string // Access Point Name
	RSSI     string
	SINR     string
	RSCP     string
	ECIO     string // EC/IO
	Bands    []string
}

func NewModemResponse(m *modems.ModemInfo, s *sim.SimInfo, signal *modems.ModemSignal) *ModemResponse {
	return &ModemResponse{
		Operator: s.Properties.OperatorName,
		APN:      m.ThreeGpp.Eps.InitialBearer.Settings.Apn,
		Bands:    m.Generic.CurrentBands,
		RSSI:     "",
	}
}

func ModemsListHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		modemsList, err := modems.List()
		if err != nil {
			log.Errorf("Failed get list of available modems: %s", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		info := []*modems.ModemInfo{}

		for _, m := range modemsList {
			i, err := modems.Get(m)

			if err != nil {
				log.Warnf("Failed get modem '%s' info: %s", m, err)
				continue
			}
			info = append(info, i)
		}
		ctx.JSON(http.StatusOK, gin.H{"modems": info})
	}
}
func DisableModemHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		modemId := ctx.Param("modem")
		modem, err := modems.Get(modemId)
		if err != nil {
			log.Errorf("Failed get modem '%s' info: %s", modem, err)
			ctx.AbortWithStatus(http.StatusNotFound)
		}
		if err = modem.Disable(); err != nil {
			log.Errorf("Failed disable modem '%s': %s", modem, err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		ctx.JSON(http.StatusOK, gin.H{"enabled": false})
	}
}

func EnableModemHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		modemId := ctx.Param("modem")
		modem, err := modems.Get(modemId)
		if err != nil {
			log.Errorf("Failed get modem '%s' info: %s", modem, err)
			ctx.AbortWithStatus(http.StatusNotFound)
		}
		if err = modem.Enable(); err != nil {
			log.Errorf("Failed enable modem '%s': %s", modem, err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		ctx.JSON(http.StatusOK, gin.H{"enabled": true})
	}
}
func GetModemSignalHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		modemId := ctx.Param("modem")
		modem, err := modems.Get(modemId)
		if err != nil {
			log.Errorf("Failed get modem '%s' info: %s", modem, err)
			ctx.AbortWithStatus(http.StatusNotFound)
		}
		signal, err := modem.GetSignal()
		if err != nil {
			log.Errorf("Failed get signal for modem '%s': %s", modem, err)
		}
		ctx.JSON(http.StatusOK, signal)
	}
}
