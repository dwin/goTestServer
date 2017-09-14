package utils

var (
	AppHostname    = "http-test.net"
	Environment    = "development"
	StaticBasePath = ""
)

func init() {
	if Environment == "development" {
		StaticBasePath = ".."
	}
}
