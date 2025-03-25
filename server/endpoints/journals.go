package endpoints

import (
	"net/http"
	"ras/management/journals"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func JournalsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		journalName := c.Param("journal")
		switch journalName {
		case "system":
			journal, err := journals.System()
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{"journal": journal})
			return
		case "core":
			journal, err := journals.Core()
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{"journal": journal})
			return
		case "connections":
			journal, err := journals.Connections()
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{"journal": journal})
			return
		case "port-forwarding":
			journal, err := journals.PortForwarding()
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{"journal": journal})
			return
		default:
			log.Warnf("No such journal named '%s'", journalName)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "There is not such journal",
			})

		}
	}
}
