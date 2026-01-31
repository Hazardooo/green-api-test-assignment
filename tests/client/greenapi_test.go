package client

import (
	"context"
	client2 "green-api-test-assignment/internal/client"
	"green-api-test-assignment/internal/models/greenapi"
	envTest "green-api-test-assignment/tests"
	"net/http"
	"testing"
	"time"
)

var (
	httpClientMock = &http.Client{Timeout: 30 * time.Second}
	ctxMock        = context.Background()
)

func TestGreenAPIClient_GetSettings(t *testing.T) {
	type fields struct {
		httpClient *http.Client
		baseURL    string
		idInstance string
		apiToken   string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    greenapi.GetSettingsResponse
		wantErr bool
	}{
		{
			name: "GetSettings",
			fields: fields{
				httpClient: httpClientMock,
				baseURL:    client2.BaseURL,
				idInstance: envTest.IdInstanceMock,
				apiToken:   envTest.ApiTokenMock,
			},
			args: args{
				ctx: ctxMock,
			},
			want:    greenapi.GetSettingsResponse{},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client2.GreenAPIClient{
				HttpClient: tt.fields.httpClient,
				BaseURL:    tt.fields.baseURL,
				IdInstance: tt.fields.idInstance,
				ApiToken:   tt.fields.apiToken,
			}
			got, err := c.GetSettings(tt.args.ctx)
			if err != nil {
				t.Log(err)
			}
			t.Logf("GetSettings() got = %v", got)
		})
	}
}

func TestGreenAPIClient_GetStateInstance(t *testing.T) {
	type fields struct {
		httpClient *http.Client
		baseURL    string
		idInstance string
		apiToken   string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    greenapi.GetStateInstanceResponse
		wantErr bool
	}{
		{
			name: "GetStateInstance",
			fields: fields{
				httpClient: httpClientMock,
				baseURL:    client2.BaseURL,
				idInstance: envTest.IdInstanceMock,
				apiToken:   envTest.ApiTokenMock,
			},
			args: args{
				ctx: ctxMock,
			},
			want:    greenapi.GetStateInstanceResponse{},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client2.GreenAPIClient{
				HttpClient: tt.fields.httpClient,
				BaseURL:    tt.fields.baseURL,
				IdInstance: tt.fields.idInstance,
				ApiToken:   tt.fields.apiToken,
			}
			got, err := c.GetStateInstance(tt.args.ctx)
			if err != nil {
				t.Log(err)
			}
			t.Logf("GetStateInstance() got = %v", got)
		})
	}
}

func TestGreenAPIClient_SendFileByUrl(t *testing.T) {
	type fields struct {
		httpClient *http.Client
		baseURL    string
		idInstance string
		apiToken   string
	}
	type args struct {
		ctx context.Context
		req greenapi.SendFileRequest
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse greenapi.SendMessageResponse
		wantErr      bool
	}{
		{
			name: "",
			fields: fields{
				httpClient: httpClientMock,
				baseURL:    client2.BaseURL,
				idInstance: envTest.IdInstanceMock,
				apiToken:   envTest.ApiTokenMock,
			},
			args: args{
				ctx: ctxMock,
				req: greenapi.SendFileRequest{
					ChatIdOrNumber: envTest.TestNumberMock,
					UrlFile:        "https://media.foma.ru/2020/12/candle-794312_1920.jpg",
					FileName:       "candle-794312_1920.jpg",
					Caption:        "свечки",
				},
			},
			wantResponse: greenapi.SendMessageResponse{},
			wantErr:      false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client2.GreenAPIClient{
				HttpClient: tt.fields.httpClient,
				BaseURL:    tt.fields.baseURL,
				IdInstance: tt.fields.idInstance,
				ApiToken:   tt.fields.apiToken,
			}
			got, err := c.SendFileByUrl(tt.args.ctx, tt.args.req)
			if err != nil {
				t.Log(err)
			}
			t.Logf("SendFileByUrl() got = %v", got)
		})
	}
}

func TestGreenAPIClient_SendMessage(t *testing.T) {
	type fields struct {
		httpClient *http.Client
		baseURL    string
		idInstance string
		apiToken   string
	}
	type args struct {
		ctx context.Context
		req greenapi.SendMessageRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    greenapi.SendMessageResponse
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				httpClient: httpClientMock,
				baseURL:    client2.BaseURL,
				idInstance: envTest.IdInstanceMock,
				apiToken:   envTest.ApiTokenMock,
			},
			args: args{
				ctx: ctxMock,
				req: greenapi.SendMessageRequest{
					ChatIdOrNumber: envTest.TestNumberMock,
					Message:        "...",
				},
			},
			want:    greenapi.SendMessageResponse{},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client2.GreenAPIClient{
				HttpClient: tt.fields.httpClient,
				BaseURL:    tt.fields.baseURL,
				IdInstance: tt.fields.idInstance,
				ApiToken:   tt.fields.apiToken,
			}
			got, err := c.SendMessage(tt.args.ctx, tt.args.req)
			if err != nil {
				t.Log(err)
			}
			t.Logf("SendMessage() got = %v", got)
		})
	}
}
