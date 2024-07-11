package proxy

import (
	"errors"
	"http-proxy-server/internal/request"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func ProxyRequest(data *request.RequestData) ([]byte, int, map[string]string, error) {
	if data.Method == "" {
		err := errors.New("HTTP method is required")
		log.Println(err)
		return nil, 0, nil, err
	}

	if data.URL == "" {
		err := errors.New("URL is required")
		log.Println(err)
		return nil, 0, nil, err
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
		err := errors.New("invalid HTTP method")
		log.Println(err)
		return nil, 0, nil, err
	}

	req, err := http.NewRequest(data.Method, data.URL, nil)
	if err != nil {
		log.Printf("Error creating new request: %v", err)
		return nil, 0, nil, err
	}

	for header, value := range data.Headers {
		req.Header.Set(header, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error performing request: %v", err)
		return nil, 0, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return nil, 0, nil, err
	}

	headers := make(map[string]string)
	for key, values := range resp.Header {
		headers[key] = strings.Join(values, ",")
	}

	return body, resp.StatusCode, headers, nil
}
