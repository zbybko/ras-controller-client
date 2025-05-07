package endpoints

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func NetloadHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		interfaces, _ := net.Interfaces()

		netload := Netload{}
		for _, iface := range interfaces {
			load, err := getNetload(iface)
			if err != nil {
				log.Errorf("Failed to get netload for interface %s: %s", iface.Name, err)
				continue
			}
			netload.Append(load)
		}

		ctx.JSON(http.StatusOK, netload)
	}
}

// TODO: unduplicate reading value from file
func getNetload(i net.Interface) (*Netload, error) {
	result := &Netload{}
	data, err := os.ReadFile(fmt.Sprintf("/sys/class/net/%s/statistics/rx_bytes", i.Name))
	if err != nil {
		return nil, err
	}
	result.RxBytes, err = strconv.Atoi(
		strings.TrimSpace(
			string(data),
		),
	)
	if err != nil {
		return nil, err
	}
	data, err = os.ReadFile(fmt.Sprintf("/sys/class/net/%s/statistics/tx_bytes", i.Name))
	if err != nil {
		return nil, err
	}
	result.TxBytes, err = strconv.Atoi(
		strings.TrimSpace(
			string(data),
		),
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type Netload struct {
	RxBytes int `json:"rxBytes"`
	TxBytes int `json:"txBytes"`
}

func (n *Netload) Append(other *Netload) {
	n.RxBytes += other.RxBytes
	n.TxBytes += other.TxBytes
}
