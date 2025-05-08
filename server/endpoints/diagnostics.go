package endpoints

import (
	"errors"
	"net/http"
	"ras/management/diagnostics"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func PingHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		addr := ctx.Param("address")
		output, err := diagnostics.Ping(addr)
		if err != nil {
			log.Errorf("Failed ping '%s': %s", addr, err)
			ctx.JSON(http.StatusBadRequest, gin.H{"output": output, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"output": output})
	}
}

func NslookupHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		addr := ctx.Param("address")
		output, err := diagnostics.Nslookup(addr)
		if err != nil {
			log.Errorf("Failed nslookup '%s': %s", addr, err)
			ctx.JSON(http.StatusBadRequest, gin.H{"output": output, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"output": output})
	}
}
func DefaultPingAddressHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defaultAddress := diagnostics.GetDefaultPingAddress()
		ctx.JSON(http.StatusOK, gin.H{"defaultAddress": defaultAddress})
	}
}

func TracerouteHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		addr := ctx.Param("address")
		output, err := diagnostics.Traceroute(addr)
		if err != nil && errors.Is(err, diagnostics.ErrEmptyAddress) {
			log.Errorf("No address specified for traceroute %s", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else if err != nil {
			log.Errorf("Failed traceroute: %s", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return

		}

		ctx.JSON(http.StatusOK, gin.H{"output": output})
	}
}
