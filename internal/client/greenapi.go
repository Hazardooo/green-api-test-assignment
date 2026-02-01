package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"green-api-test-assignment/internal/models/greenapi"
	utils "green-api-test-assignment/internal/utils"
)

const (
	BaseURL = "https://1103.api.green-api.com/waInstance"
)

type GreenAPIClient struct {
	HttpClient *http.Client
	BaseURL    string
	IdInstance string
	ApiToken   string
}

func New(idInstance, apiToken string) *GreenAPIClient {
	return &GreenAPIClient{
		HttpClient: &http.Client{Timeout: 30 * time.Second},
		BaseURL:    utils.GetURLClient(BaseURL, idInstance),
		IdInstance: idInstance,
		ApiToken:   apiToken,
	}
}

func (c *GreenAPIClient) request(ctx context.Context, method, endpoint string, payload, result any) error {

	url := utils.GetEndPointPOSTURL(c.BaseURL, endpoint, c.ApiToken)

	body, err := utils.ToReader(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}
	return nil
}

func (c *GreenAPIClient) GetStateInstance(ctx context.Context) (greenapi.GetStateInstanceResponse, error) {
	var res greenapi.GetStateInstanceResponse
	err := c.request(ctx, http.MethodGet, "getStateInstance", nil, &res)
	return res, err
}

func (c *GreenAPIClient) GetSettings(ctx context.Context) (greenapi.GetSettingsResponse, error) {
	var res greenapi.GetSettingsResponse
	err := c.request(ctx, http.MethodGet, "getSettings", nil, &res)
	return res, err
}

func (c *GreenAPIClient) SendMessage(ctx context.Context, req greenapi.SendMessageRequest) (greenapi.SendMessageResponse, error) {
	req.ChatIdOrNumber = fmt.Sprintf("%s@c.us", req.ChatIdOrNumber)
	var res greenapi.SendMessageResponse
	err := c.request(ctx, http.MethodPost, "sendMessage", req, &res)
	return res, err
}

func (c *GreenAPIClient) SendFileByUrl(ctx context.Context, req greenapi.SendFileRequest) (greenapi.SendMessageResponse, error) {
	req.ChatIdOrNumber = fmt.Sprintf("%s@c.us", req.ChatIdOrNumber)
	var res greenapi.SendMessageResponse
	err := c.request(ctx, http.MethodPost, "sendFileByUrl", req, &res)
	return res, err
}
