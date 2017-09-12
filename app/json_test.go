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

func TestGetIPJSON(t *testing.T) {
	r := gofight.New()
	//r.SetHeader(gofight.H{"HTTP_CLIENT_IP": "0.0.0.0"})
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
