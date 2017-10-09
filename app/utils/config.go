package utils

import "time"

var (
	AppHostname    = "http-test.net"
	Environment    = "development"
	StaticBasePath = ""
	StartTime      = time.Now()
	DBuser         = "postgres"
	DBpass         = ""
	DBname         = "gotestserver"
	DBhost         = "localhost"
)

func init() {
	if Environment == "development" {
		StaticBasePath = ".."
	} else {
		StaticBasePath = ""
		DBhost = "postgres"
	}
}
