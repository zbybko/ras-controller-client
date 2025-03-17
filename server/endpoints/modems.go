package endpoints

import (
	"net/http"
	"ras/management/modems"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func ModemsListHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		modemsList, err := modems.GetList()
		if err != nil {
			log.Errorf("Failed get list of available modems: %s", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		info := []*modems.ModemInfo{}

		for _, m := range modemsList {
			i, err := modems.GetInfo(m)

			if err != nil {
				log.Warnf("Failed get modem '%s' info: %s", m, err)
				continue
			}
			info = append(info, i)
		}
		ctx.JSON(http.StatusOK, gin.H{"modems": info})
	}
}
