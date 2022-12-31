package cmd

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

func getBody(url string) ([]byte, error) {
	c := http.Client{
		Transport: &http.Transport{
			Proxy:           http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	res, err := c.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Cant connect to %s error %s", Host, err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http returned witn status code %s", res.Status)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return b, nil
}
