package httputils

import (
	"net/http"
	"crypto/tls"
	"strings"
	"time"
)

func ProbeURL(target string, timeoutSeconds int) (string, string) {
	target = strings.TrimSpace(target)

	if strings.HasPrefix(target, "http://") || strings.HasPrefix(target, "https://") {
		if isReachable(target, timeoutSeconds) {
			return target, ""
		}
		return "", target
	}

	httpsURL := "https://" + target
	if isReachable(httpsURL, timeoutSeconds) {
		return httpsURL, ""
	}

	httpURL := "http://" + target
	if isReachable(httpURL, timeoutSeconds) {
		return httpURL, ""
	}
	
	return "", target
}

func isReachable(url string, timeout int) bool {
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Head(url)
	if err != nil {
		resp, err = client.Get(url)
		if err != nil {
			return false
		}
	}

	defer resp.Body.Close()

	if resp.StatusCode < 400 {
		return true
	}
	
	return false
}
