package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

func TestGetCookiesJSON(t *testing.T) {
	r := gofight.New()
	r.GET("/json/cookies").SetDebug(true).Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, 200, resp.Code)
		assert.Contains(t, "cookies", r.Body)
	})
}
func TestDeleteCookiesJSON(t *testing.T) {
	r := gofight.New()
	r.GET("/json/cookies/delete").SetDebug(true).Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, 302, resp.Code)
	})
}
func TestGetJSON(t *testing.T) {
	r := gofight.New()
	r.GET("/json").SetDebug(true).Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, 200, resp.Code)
		assert.Contains(t, "origin-ip", r.Body)
	})
}
func TestGetDateTimeJSON(t *testing.T) {
	r := gofight.New()
	r.GET("/json/datetime").SetDebug(true).Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, 200, resp.Code)
		assert.Contains(t, "ANSIC", r.Body)
		assert.Contains(t, "RFC3339", r.Body)
	})
}

func TestGetIPJSON(t *testing.T) {
	r := gofight.New()
	r.GET("/json/ip").SetDebug(true).Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, 200, resp.Code)
		assert.Contains(t, "origin-ip", r.Body)
	})
}
func TestGetHeadersJSON(t *testing.T) {
	randNum := rand.New(rand.NewSource(99)).Int31
	testString := fmt.Sprintf("%v", randNum())
	r := gofight.New()
	r.SetHeader(gofight.H{"Test String": testString})
	r.GET("/json/headers").SetDebug(true).Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, 200, resp.Code)
		assert.Contains(t, testString, r.Body)
	})
}
func TestGetUserAgentJSON(t *testing.T) {
	r := gofight.New()
	r.GET("/json/user-agent").SetDebug(true).Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, 200, resp.Code)
		assert.Contains(t, "Gofight-client/1.0", r.Body)
	})
}
func TestGetStatusJSON(t *testing.T) {
	r := gofight.New()
	r.GET("/json/status").SetDebug(true).Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, 200, resp.Code)
		assert.Contains(t, "ok", r.Body)
	})
}
func TestGetUUIDv1JSON(t *testing.T) {
	r := gofight.New()
	r.GET("/json/uuid/1").SetDebug(true).Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, 200, resp.Code)
		assert.Contains(t, "UUID-v1", r.Body)
	})
}
func TestGetUUIDv4JSON(t *testing.T) {
	r := gofight.New()
	r.GET("/json/uuid/4").SetDebug(true).Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, 200, resp.Code)
		assert.Contains(t, "UUID-v4", r.Body)
	})
}
func TestGetUUIDInvalidJSON(t *testing.T) {
	r := gofight.New()
	r.GET("/json/uuid/10").SetDebug(true).Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, 400, resp.Code)
	})
}
func TestSetCookieJSON(t *testing.T) {
	r := gofight.New()
	r.GET("/json/cookies/set?name=Jane&value=Doe").SetDebug(true).Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, 302, resp.Code)
	})
	r.GET("/json/cookies").SetDebug(true).Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, 200, resp.Code)
		assert.Contains(t, "cookies", r.Body)
	})
}
func TestPostBase32JSON(t *testing.T) {
	r := gofight.New()
	r.POST("/json/base32").
		SetDebug(true).
		Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
			assert.Equal(t, 200, resp.Code)
			assert.Contains(t, "base32", r.Body)
		})
}
func TestPostBase64JSON(t *testing.T) {
	r := gofight.New()
	r.POST("/json/base64").
		SetDebug(true).
		Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
			assert.Equal(t, 200, resp.Code)
			assert.Contains(t, "base64", r.Body)
		})
}
func TestPostBlake2bJSON(t *testing.T) {
	r := gofight.New()
	r.POST("/json/blake2b/512").
		SetDebug(true).
		Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
			assert.Equal(t, 200, resp.Code)
			assert.Contains(t, "786a02f742015903c6c6fd852552d272912f4740e15847618a86e217f71f5419d25e1031afee585313896444934eb04b903a685b1448b755d56f701afe9be2ce", r.Body)
		})
	r.POST("/json/blake2b/256").
		SetDebug(true).
		Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
			assert.Equal(t, 200, resp.Code)
			assert.Contains(t, "0e5751c026e543b2e8ab2eb06099daa1d1e5df47778f7787faab45cdf12fe3a8", r.Body)
		})
}
func TestPostMD5JSON(t *testing.T) {
	r := gofight.New()
	r.POST("/json/md5").
		SetDebug(true).
		Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
			assert.Equal(t, 200, resp.Code)
			assert.Contains(t, "md5", r.Body)
		})
}
func TestPostSHA1JSON(t *testing.T) {
	r := gofight.New()
	r.POST("/json/sha1").
		SetDebug(true).
		Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
			assert.Equal(t, 200, resp.Code)
			assert.Contains(t, "da39a3ee5e6b4b0d3255bfef95601890afd80709", r.Body)
		})
}
func TestPostSHA256JSON(t *testing.T) {
	r := gofight.New()
	r.POST("/json/sha256").
		SetDebug(true).
		Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
			assert.Equal(t, 200, resp.Code)
			assert.Contains(t, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", r.Body)
		})
}
func TestPostSHA3JSON(t *testing.T) {
	r := gofight.New()
	r.POST("/json/sha3/512").
		SetDebug(true).
		Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
			assert.Equal(t, 200, resp.Code)
			assert.Contains(t, "sha3-512", r.Body)
		})
	r.POST("/json/sha3/384").
		SetDebug(true).
		Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
			assert.Equal(t, 200, resp.Code)
			assert.Contains(t, "sha3-384", r.Body)
		})
	r.POST("/json/sha3/256").
		SetDebug(true).
		Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
			assert.Equal(t, 200, resp.Code)
			assert.Contains(t, "sha3-256", r.Body)
		})
	r.POST("/json/sha3/224").
		SetDebug(true).
		Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
			assert.Equal(t, 200, resp.Code)
			assert.Contains(t, "sha3-224", r.Body)
		})
}
