package internal

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func GetHttpProxyClient() *http.Client {
	proxy := os.Getenv("PROXY")

	parts := strings.Split(proxy, ":")
	if len(parts) != 4 {
		fmt.Println("Invalid proxy string format")
		os.Exit(8)
	}

	ip := parts[0]
	port := parts[1]
	username := parts[2]
	password := parts[3]

	proxyURL := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", ip, port),
		User:   url.UserPassword(username, password),
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	return &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}
}
