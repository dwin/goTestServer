package controller

import (
	"strconv"

	u "github.com/dwin/goTestServer/app/utils"
	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	statusReq, err := strconv.Atoi(c.Param("status"))
	if err != nil {
		u.Log.Error().Err(err).Msg("Error Parsing int in GetStatus")
		c.AbortWithError(500, err)
		return
	}
	if statusReq < 0 || statusReq > 999 {
		u.Log.Warn().Int("Status Request", statusReq).Msg("Status Request outside acceptable range")
		c.JSON(400, gin.H{
			"error": "Error request must be in range 0-999",
		})
		return
	}
	c.Status(statusReq)
}

// GetRedirects returns redirect code time specified by client in parameter
func GetRedirects(c *gin.Context) {
	redirs, err := strconv.Atoi(c.Param("redirects"))
	if err != nil {
		u.Log.Error().Err(err).Msg("Error parsing int in GetRedirects")
		c.AbortWithError(500, err)
		return
	}
	for i := 0; i < redirs; i++ {
		c.Redirect(302, "/json/get")
	}
	return
}
