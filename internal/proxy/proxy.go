package proxy

import (
	"errors"
	"http-proxy-server/internal/request"
	"io/ioutil"
	"net/http"
	"strings"
)

func ProxyRequest(data *request.RequestData) ([]byte, int, map[string]string, error) {
	if data.Method == "" {
		return nil, 0, nil, errors.New("HTTP method is required")
	}

	if data.URL == "" {
		return nil, 0, nil, errors.New("URL is required")
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
		return nil, 0, nil, errors.New("invalid HTTP method")
	}

	req, err := http.NewRequest(data.Method, data.URL, nil)
	if err != nil {
		return nil, 0, nil, err
	}

	for header, value := range data.Headers {
		req.Header.Set(header, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, nil, err
	}

	headers := make(map[string]string)
	for key, values := range resp.Header {
		headers[key] = strings.Join(values, ",")
	}

	return body, resp.StatusCode, headers, nil
}
