package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"input-api/models"
)

func GetNodeApiUrl() string {
	nodeApiUrl := os.Getenv("NODE_API_URL")

	if nodeApiUrl == "" {
		nodeApiUrl = "http://localhost:4000/statistics"
	}
	return nodeApiUrl
}

func SendQRToNode(result *models.QrResponse) (map[string]interface{}, error) {

	body, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(GetNodeApiUrl(), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("node API returned status: %d", response.StatusCode)
	}

	var statistics map[string]interface{}

	if err := json.NewDecoder(response.Body).Decode(&statistics); err != nil {
		return nil, err
	}

	return statistics, nil
}
