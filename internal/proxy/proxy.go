package proxy

import (
	"errors"
	"http-proxy-server/internal/request"
	"io/ioutil"
	"net/http"
	"strings"
)

func ProxyRequest(data *request.RequestData) ([]byte, error) {
	if data.Method == "" {
		return nil, errors.New("HTTP method is required")
	}

	if data.URL == "" {
		return nil, errors.New("URL is required")
	}

	validMethods := map[string]bool{
		"GET":     true,
		"POST":    true,
		"PUT":     true,
		"DELETE":  true,
		"PATCH":   true,
		"HEAD":    true,
		"OPTIONS": true,
	}

	if !validMethods[strings.ToUpper(data.Method)] {
		return nil, errors.New("invalid HTTP method")
	}
	req, err := http.NewRequest(data.Method, data.URL, nil)
	if err != nil {
		return nil, err
	}

	for headers, value := range data.Headers {
		req.Header.Set(headers, value)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
