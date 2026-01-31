package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"green-api-test-assignment/internal/models/responses"
	"green-api-test-assignment/internal/utils"
	"io"
	"net/http"
	"time"
)

const (
	BaseUrl = "https://1103.api.green-api.com/waInstance"
)

type GreenApiClient struct {
	HttpClient *http.Client
	BaseURL    string
	IdInstance string
	ApiToken   string
}

func NewGreenClient(idInstance, apiToken string) *GreenApiClient {
	return &GreenApiClient{
		HttpClient: &http.Client{Timeout: 30 * time.Second},
		BaseURL:    utils.GetUrlClient(BaseUrl, idInstance),
		IdInstance: idInstance,
		ApiToken:   apiToken,
	}
}

func (client *GreenApiClient) GetStateInstance(ctx context.Context) (err error, clientResponse responses.GetStateInstanceResponse) {
	url := fmt.Sprintf("%s/getStateInstance/%s", client.BaseURL, client.ApiToken)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		err = errors.New(fmt.Sprintf("Ошибка в создании запроса: %s", err))
		return err, clientResponse
	}
	resp, err := client.HttpClient.Do(req)
	if err != nil {
		err = errors.New(fmt.Sprintf("Ошибка отправки: %s", err))
		return err, clientResponse
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("Неожиданный статус %d", resp.StatusCode))
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err = errors.New(fmt.Sprintf("Ошибка чтения ответа: %s", err))
		return
	}
	err = json.Unmarshal(body, &clientResponse)
	if err != nil {
		return
	}
	return
}

func (client *GreenApiClient) GetSettings(ctx context.Context) (err error, clientResponse responses.GetSettingsResponse) {
	url := fmt.Sprintf("%s/getSettings/%s", client.BaseURL, client.ApiToken)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		err = errors.New(fmt.Sprintf("Ошибка в создании запроса: %s", err))
		return err, clientResponse
	}
	resp, err := client.HttpClient.Do(req)
	if err != nil {
		err = errors.New(fmt.Sprintf("Ошибка отправки: %s", err))
		return err, clientResponse
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("Неожиданный статус %d", resp.StatusCode))
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err = errors.New(fmt.Sprintf("Ошибка чтения ответа: %s", err))
		return
	}
	err = json.Unmarshal(body, &clientResponse)
	if err != nil {
		return
	}
	return
}
