package main

import (
	"fmt"
	"http-proxy-server/internal/proxy"
	"http-proxy-server/internal/request"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	requestMap  sync.Map
	responseMap sync.Map
)

func main() {
	http.HandleFunc("/", handleFunc)
	fmt.Println("Starting server on :8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		return
	}
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	requestData, err := request.ExtractRequestData(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	requestID := fmt.Sprintf("%d", time.Now().UnixNano())
	requestMap.Store(requestID, requestData)

	body, err := proxy.ProxyRequest(requestData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	responseMap.Store(requestID, body)

	log.Println("response map ", responseMap)
	log.Println("request map ", requestMap)

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
