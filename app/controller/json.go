package controller

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"

	u "github.com/dwin/goTestServer/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/minio/blake2b-simd"
	"github.com/minio/sha256-simd"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/sha3"
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
	if len(c.ClientIP()) < 6 {
		// Log Failure
		u.Log.Error().Msg("Error: Could not obtain client IP")
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
	u.Log.Error().Str("IP", c.ClientIP()).Msg("Could not read user-agent from client")
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
		c.String(400, "Must request version. ex. /json/uuid/4 or /json/uuid/1")
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
func PostBase32JSON(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		u.Log.Error().Err(err).Msg("Error reading request body for PostBase32JSON")
		c.Status(500)
		return
	}
	base32str := base32.StdEncoding.EncodeToString(body)
	c.IndentedJSON(200, gin.H{
		"base32": base32str,
	})
	return
}
func PostBase64JSON(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		u.Log.Error().Err(err).Msg("Error reading request body for PostBase64JSON")
		c.Status(500)
		return
	}
	base64str := base64.StdEncoding.EncodeToString(body)
	c.IndentedJSON(200, gin.H{
		"base64": base64str,
	})
	return
}
func PostBlake2bJSON(c *gin.Context) {
	switch ver := c.Param("size"); ver {
	case "512":
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			u.Log.Error().Err(err).Msg("Error reading request body for PostBlake2b512JSON")
			c.Status(500)
			return
		}
		hash := blake2b.Sum512(body)
		b2bHex := hex.EncodeToString(hash[:])
		c.IndentedJSON(200, gin.H{
			"blake2b-512": b2bHex,
		})
		return
	case "256":
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			u.Log.Error().Err(err).Msg("Error reading request body for PostBlake2b256JSON")
			c.Status(500)
			return
		}
		hash := blake2b.Sum256(body)
		b2bHex := hex.EncodeToString(hash[:])
		c.IndentedJSON(200, gin.H{
			"blake2b-256": b2bHex,
		})
		return
	default:
		c.String(400, "Unsupported version. ex. /json/blake2b/512 or /json/blake2b/256")
		return
	}
}
func PostMD5JSON(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		u.Log.Error().Err(err).Msg("Error reading request body for PostMD5JSON")
		c.Status(500)
		return
	}
	hash := md5.Sum(body)
	md5hex := hex.EncodeToString(hash[:])
	c.IndentedJSON(200, gin.H{
		"md5": md5hex,
	})
	return
}
func PostSHA1JSON(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		u.Log.Error().Err(err).Msg("Error reading request body for PostSHA1JSON")
		c.Status(500)
		return
	}
	hash := sha1.Sum(body)
	sha1hex := hex.EncodeToString(hash[:])
	c.IndentedJSON(200, gin.H{
		"sha1": sha1hex,
	})
	return
}
func PostSHA256JSON(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		u.Log.Error().Err(err).Msg("Error reading request body for PostSHA256JSON")
		c.Status(500)
		return
	}
	hash := sha256.Sum256(body)
	sha256hex := hex.EncodeToString(hash[:])
	c.IndentedJSON(200, gin.H{
		"sha2-256": sha256hex,
	})
	return
}
func PostSHA3JSON(c *gin.Context) {
	switch ver := c.Param("size"); ver {
	case "512":
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			u.Log.Error().Err(err).Msg("Error reading request body for PostSHA3-512JSON")
			c.Status(500)
			return
		}
		hash := sha3.Sum512(body)
		sha3_512hex := hex.EncodeToString(hash[:])
		c.IndentedJSON(200, gin.H{
			"sha3-512": sha3_512hex,
		})
		return
	case "384":
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			u.Log.Error().Err(err).Msg("Error reading request body for PostSHA3-384JSON")
			c.Status(500)
			return
		}
		hash := sha3.Sum384(body)
		sha3_384hex := hex.EncodeToString(hash[:])
		c.IndentedJSON(200, gin.H{
			"sha3-384": sha3_384hex,
		})
		return
	case "256":
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			u.Log.Error().Err(err).Msg("Error reading request body for PostSHA3-256JSON")
			c.Status(500)
			return
		}
		hash := sha3.Sum256(body)
		sha3_256hex := hex.EncodeToString(hash[:])
		c.IndentedJSON(200, gin.H{
			"sha3-256": sha3_256hex,
		})
		return
	case "224":
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			u.Log.Error().Err(err).Msg("Error reading request body for PostSHA3-224JSON")
			c.Status(500)
			return
		}
		hash := sha3.Sum224(body)
		sha3_224hex := hex.EncodeToString(hash[:])
		c.IndentedJSON(200, gin.H{
			"sha3-224": sha3_224hex,
		})
		return
	default:
		c.String(400, "Unsupported version. ex. /json/sha3/512 , /json/sha3/384 , /json/sha3/256 , /json/sha3/224")
		return
	}
}
