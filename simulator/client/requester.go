package client

import (
	"Regret/domain"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type simulatorClient struct {
	baseURL string
}

func (s *simulatorClient) SendMessage(message *domain.Message) *domain.Carrier {
	var carrier domain.Carrier
	url := s.baseURL + "/carrier"

	bodyBytes, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Error marshaling message to JSON: %v", err)
	}

	bodyReader := bytes.NewReader(bodyBytes)
	carrierRequest, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		log.Fatalf("Error while getting the carrier for request : %v", err)
	}

	carrierRequest.Header.Set("Content-Type", "application/json")
	SetTickHeader(carrierRequest)

	resp, requestErr := http.DefaultClient.Do(carrierRequest)
	if requestErr != nil {
		log.Fatalf("Error while executing the carrier request: %v", requestErr)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Error while closing the body")
		}
	}(resp.Body)

	err = json.NewDecoder(resp.Body).Decode(&carrier)
	if err != nil {
		log.Fatalf("Error while decoding carrier response: %v", err)
	}

	return &carrier
}

type SimulatorClient interface {
	SendMessage(message *domain.Message) *domain.Carrier
}

func NewSimulatorClient(baseURL string) SimulatorClient {
	return &simulatorClient{
		baseURL: baseURL,
	}
}
