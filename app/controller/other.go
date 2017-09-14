package controller

import (
	"fmt"
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
	redirs, err := strconv.Atoi(c.Param("num"))
	if err != nil {
		u.Log.Error().Err(err).Msg("Error parsing int in GetRedirects")
		c.AbortWithError(500, err)
		return
	}
	if redirs < 1 {
		fmt.Println("less than 1")
		c.Redirect(302, "/")
	}

	c.Redirect(302, fmt.Sprintf("/redirect/%v", redirs-1))
	return
}
