package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	greenapirequests "green-api-test-assignment/internal/models/greenapi"
	"io"
	"net/http"
	"time"

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
	url := utils.GetEndPointURL(c.BaseURL, c.IdInstance, c.ApiToken, endpoint)
	var bodyReader io.Reader
	if payload != nil {
		body, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("marshal payload: %w", err)
		}
		bodyReader = bytes.NewReader(body)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
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
		return fmt.Errorf("unexpected status %d", resp.StatusCode)
	}
	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
	if result != nil {
		json.Unmarshal(bodyBytes, &result)
		fmt.Println(result)
	}
	return nil
}

func (c *GreenAPIClient) GetStateInstance(ctx context.Context) (greenapirequests.GetStateInstanceResponse, error) {
	var res greenapirequests.GetStateInstanceResponse
	err := c.request(ctx, http.MethodGet, "getStateInstance", nil, &res)
	return res, err
}

func (c *GreenAPIClient) GetSettings(ctx context.Context) (greenapirequests.GetSettingsResponse, error) {
	var res greenapirequests.GetSettingsResponse
	err := c.request(ctx, http.MethodGet, "getSettings", nil, &res)
	return res, err
}

func (c *GreenAPIClient) SendMessage(ctx context.Context, req greenapirequests.SendMessageRequest) (response greenapirequests.SendMessageResponse, err error) {
	req.ChatIdOrNumber = fmt.Sprintf("%s@c.us", req.ChatIdOrNumber)
	err = c.request(ctx, http.MethodPost, "sendMessage", req, &response)
	return
}

func (c *GreenAPIClient) SendFileByUrl(ctx context.Context, req greenapirequests.SendFileRequest) (response greenapirequests.SendMessageResponse, err error) {
	req.ChatIdOrNumber = fmt.Sprintf("%s@c.us", req.ChatIdOrNumber)
	err = c.request(ctx, http.MethodPost, "sendFileByUrl", req, &response)
	return
}
