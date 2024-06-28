package request

import (
	"encoding/json"
	"net/http"
)

type RequestData struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

func ExtractRequestData(req *http.Request) (*RequestData, error) {
	var requestData RequestData
	err := json.NewDecoder(req.Body).Decode(&requestData)

	if err != nil {
		return nil, err
	}
	return &requestData, nil
}
