package libhoney

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	markerAPITimeout = 3 * time.Second
)

type MarkerClientConfig struct {
	APIKey string
	APIHost string
	Dataset string
}

type markerClient struct {
	apiKey string
	apiHost string
	dataset string
	client http.Client
}

func NewMarkerClient(cfg MarkerClientConfig) *markerClient {
	if cfg.APIHost == "" {
		cfg.APIHost = defaultAPIHost
	}
	return &markerClient{
		client: http.Client{
			Timeout: markerAPITimeout,
		},
		apiKey: cfg.APIKey,
		apiHost: cfg.APIHost,
		dataset: cfg.Dataset,
	}
}

func (m *markerClient) SendMarker(marker *Marker) (string, error) {
	data, jsonErr := json.Marshal(marker)
	if jsonErr != nil {
		return "", jsonErr
	}

	endpoint := m.apiHost + "1/markers/" + m.dataset
	req, reqErr := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(data))
	if reqErr != nil {
		return "", reqErr
	}
	req.Header.Set("X-Honeycomb-Team", m.apiKey)

	resp, respErr := m.client.Do(req)
	if respErr != nil {
		return "", respErr
	}
	defer resp.Body.Close()

	respData, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return "", readErr
	}

	// this is obviously a minimal response
	body := struct{
		ID string `json:"id"`
	}{}
	if umErr := json.Unmarshal(respData, &body); umErr != nil {
		return "", umErr
	}

	return body.ID, nil
}

func (m *markerClient) DeleteMarker(id string) error {
	endpoint := m.apiHost + "1/markers/" + m.dataset + "/" + id
	req, reqErr := http.NewRequest(http.MethodDelete, endpoint, http.NoBody)
	if reqErr != nil {
		return reqErr
	}
	req.Header.Set("X-Honeycomb-Team", m.apiKey)

	_, respErr := m.client.Do(req)

	return respErr
}
