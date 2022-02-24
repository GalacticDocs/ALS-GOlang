package als_golang

import (
	"net/http"
	"time"
)

var transport = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}

var client = &http.Client{
	Transport:	transport,
}