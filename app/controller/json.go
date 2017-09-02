package controller

import (
	"fmt"

	u "github.com/dwin/goTestServer/app/utils"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// The request responds to a url matching:  /delete?name=Jane
func DeleteCookieJSON(c *gin.Context) {
	name := c.Query("name")
	c.SetCookie(name, "", 1, "/json/cookies", u.AppHostname, false, false)
	c.Redirect(302, "/json/cookies")
	return
}
func GetCookiesJSON(c *gin.Context) {
	cookies := c.Request.Header.Get("Cookie")
	c.IndentedJSON(200, gin.H{
		"cookies": cookies,
	})
	return
}

// GetJSON returns JSON with GET data
func GetJSON(c *gin.Context) {
	h := make(map[string]string)
	for k, v := range c.Request.Header {
		h[k] = fmt.Sprintf("%v", v)
	}
	c.IndentedJSON(200, gin.H{
		"args":      "",
		"headers":   h,
		"origin-ip": fmt.Sprintf("%s", c.ClientIP()),
		"uri":       c.Request.RequestURI,
	})
	return
}

// GetIPJSON returns JSON with IP address of client
func GetIPJSON(c *gin.Context) {
	if len(c.ClientIP()) < 7 {
		// Log Failure
		u.Log.Error().Msg("Error: Could not obtain client IP")
		c.String(500, "Error: Unable to obtain client IP")
		return
	}
	c.IndentedJSON(200, gin.H{
		"origin-ip": c.ClientIP(),
	})
	return
}

// GetHeadersJSON returns JSON with headers received from client
func GetHeadersJSON(c *gin.Context) {
	h := make(map[string]string)
	for k, v := range c.Request.Header {
		h[k] = fmt.Sprintf("%v", v)
	}
	c.IndentedJSON(200, h)
	return
}

// GetUserAgentJSON returns JSON with user-agent received from client
func GetUserAgentJSON(c *gin.Context) {
	if ua := c.Request.Header.Get("User-Agent"); ua != "" {
		c.IndentedJSON(200, gin.H{
			"User-Agent": ua,
		})
		return
	}
	u.Log.Error().Str("IP", c.Request.RemoteAddr).Msg("Could not read user-agent from client")
	c.IndentedJSON(400, gin.H{
		"error": "Could not read user-agent from client",
	})
	return
}

// GetUUIDJSON returns JSON with UUID of version specified in parameter
func GetUUIDJSON(c *gin.Context) {
	switch ver := c.Param("version"); ver {
	case "1":
		c.IndentedJSON(200, gin.H{
			"UUID-v" + ver: uuid.NewV1(),
		})
		return
	case "4":
		c.IndentedJSON(200, gin.H{
			"UUID-v" + ver: uuid.NewV4(),
		})
		return
	default:
		c.String(400, "Must request version. ex. /json/uuid/4")
		return
	}
}

// The request responds to a url matching:  /set?name=Jane&value=Doe
func SetCookieJSON(c *gin.Context) {
	name := c.Query("name")
	value := c.Query("value")
	c.SetCookie(name, value, 36400, "/json/cookies", u.AppHostname, false, false)
	c.Redirect(302, "/json/cookies")
	return
}
