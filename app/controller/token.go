package controller

import (
	"regexp"

	"github.com/dwin/goTestServer/app/model"
	u "github.com/dwin/goTestServer/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/ventu-io/go-shortid"
)

func GetNewToken(c *gin.Context) {
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		u.Log.Error().Err(err).Msg("Error initializing short-id token generator")
		c.AbortWithError(500, err)
		return
	}
	token, err := sid.Generate()
	if err != nil {
		u.Log.Error().Err(err).Msg("Error generating short-id token")
		c.AbortWithError(500, err)
		return
	}
	err = model.SaveNewToken(token, c.ClientIP())
	if err != nil {
		u.Log.Error().Err(err).Msg("Error saving new token to db")
		c.AbortWithError(500, err)
		return
	}
	c.HTML(200, "newToken.tmpl", gin.H{
		"token": token,
	})
}

func GetListTokenRequests(c *gin.Context) {
	// TODO: Validate token param
	t := c.Param("token")
	re := regexp.MustCompile(`([\w-]{9,11})`)
	if !re.MatchString(t) {
		c.String(400, "Invalid token")
		return
	}
	reqs, err := model.GetAllTokenRequests(t)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.HTML(200, "listTokens.tmpl", gin.H{
		"request": reqs,
	})
	return
}
