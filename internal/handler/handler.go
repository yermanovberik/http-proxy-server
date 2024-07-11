package handler

import (
	"encoding/json"
	"fmt"
	"http-proxy-server/internal/proxy"
	"http-proxy-server/internal/request"
	"http-proxy-server/internal/response"
	"net/http"
	"sync"
	"time"
)

var (
	requestMap  sync.Map
	responseMap sync.Map
)

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	requestData, err := request.ExtractRequestData(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	requestID := fmt.Sprintf("%d", time.Now().UnixNano())
	requestMap.Store(requestID, requestData)

	body, status, headers, err := proxy.ProxyRequest(requestData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseMap.Store(requestID, body)

	responseData := response.ResponseData{
		ID:      requestID,
		Status:  status,
		Headers: headers,
		Length:  len(body),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
}
