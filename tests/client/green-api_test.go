package client

import (
	"context"
	"fmt"
	client2 "green-api-test-assignment/internal/client"
	"green-api-test-assignment/internal/models/responses"
	envTest "green-api-test-assignment/tests"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var (
	baseURLMock = fmt.Sprintf("https://1103.api.green-api.com/waInstance%s", envTest.IdInstanceMock)
)

var (
	httpClientMock = http.Client{Timeout: 30 * time.Second}
)

func TestGreenApiClient_GetSettings(t *testing.T) {
	type fields struct {
		HttpClient *http.Client
		BaseURL    string
		IdInstance string
		ApiToken   string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantErr            error
		wantClientResponse responses.GetSettingsResponse
	}{
		// TODO: Add test cases.
		{
			name: "GetSettings",
			fields: fields{
				HttpClient: &httpClientMock,
				BaseURL:    baseURLMock,
				IdInstance: envTest.IdInstanceMock,
				ApiToken:   envTest.ApiTokenMock,
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr:            nil,
			wantClientResponse: responses.GetSettingsResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &client2.GreenApiClient{
				HttpClient: tt.fields.HttpClient,
				BaseURL:    tt.fields.BaseURL,
				IdInstance: tt.fields.IdInstance,
				ApiToken:   tt.fields.ApiToken,
			}
			gotErr, gotClientResponse := client.GetSettings(tt.args.ctx)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("GetSettings() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
			if !reflect.DeepEqual(gotClientResponse, tt.wantClientResponse) {
				t.Errorf("GetSettings() gotClientResponse = %v, want %v", gotClientResponse, tt.wantClientResponse)
			}
		})
	}
}

func TestGreenApiClient_GetStateInstance(t *testing.T) {
	type fields struct {
		HttpClient *http.Client
		BaseURL    string
		IdInstance string
		ApiToken   string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantErr            error
		wantClientResponse responses.GetStateInstanceResponse
	}{
		{
			name: "GetStateInstance",
			fields: fields{
				HttpClient: &httpClientMock,
				BaseURL:    baseURLMock,
				IdInstance: envTest.IdInstanceMock,
				ApiToken:   envTest.ApiTokenMock,
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: nil,
			wantClientResponse: responses.GetStateInstanceResponse{
				StateInstance: "authorized",
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &client2.GreenApiClient{
				HttpClient: tt.fields.HttpClient,
				BaseURL:    tt.fields.BaseURL,
				IdInstance: tt.fields.IdInstance,
				ApiToken:   tt.fields.ApiToken,
			}
			gotErr, gotClientResponse := client.GetStateInstance(tt.args.ctx)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("GetStateInstance() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
			if !reflect.DeepEqual(gotClientResponse, tt.wantClientResponse) {
				t.Errorf("GetStateInstance() gotClientResponse = %v, want %v", gotClientResponse, tt.wantClientResponse)
			}
		})
	}
}

func TestNewGreenClient(t *testing.T) {
	type args struct {
		idInstance string
		apiToken   string
	}
	tests := []struct {
		name string
		args args
		want *client2.GreenApiClient
	}{
		{
			name: "NewGreenClient",
			args: args{
				idInstance: envTest.IdInstanceMock,
				apiToken:   envTest.ApiTokenMock,
			},
			want: &client2.GreenApiClient{
				HttpClient: &httpClientMock,
				BaseURL:    baseURLMock,
				IdInstance: envTest.IdInstanceMock,
				ApiToken:   envTest.ApiTokenMock,
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := client2.NewGreenClient(tt.args.idInstance, tt.args.apiToken)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGreenClient() = %v, want %v", got, tt.want)
			}
			t.Log(got)
		})
	}
}
