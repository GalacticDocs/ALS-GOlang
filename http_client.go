package als_golang

import (
	"io/ioutil"
	"net/http"
	"time"
)

func GetHTTP(URL string) ([]byte, error) {
	var (
		transport = &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		}
		client = &http.Client{
			Transport: transport,
		}
	)

	resp, http_err := client.Get(URL)
	if http_err != nil {
		return nil, http_err
	}

	defer resp.Body.Close()

	body, io_err := ioutil.ReadAll(resp.Body)
	if io_err != nil {
		return nil, io_err
	}

	return body, nil
}
