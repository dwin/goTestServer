package main

import (
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
