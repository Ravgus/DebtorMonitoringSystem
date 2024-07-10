package internal

import (
	"net/http"
	"os"
)

func GetHttpClient() *http.Client {
	proxy := os.Getenv("PROXY")

	if len(proxy) != 0 {
		return GetHttpProxyClient()
	}

	return &http.Client{}
}
