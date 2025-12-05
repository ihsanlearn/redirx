package httputils

import (
	"crypto/tls"
	"net/http"
	"time"
	"github.com/ihsanlearn/redirx/internal/options"
	"net"
)

func NewScannerClient(opts *options.Options) *http.Client {
	
	disableKeepAlive := !opts.KeepAlive 

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !opts.VerifySSL},
		
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 5 * time.Second,
		}).DialContext,

		DisableKeepAlives: disableKeepAlive, 
		
		MaxIdleConns:        1000,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     4 * time.Second,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   time.Duration(opts.Timeout) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}