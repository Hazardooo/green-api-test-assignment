package utils

import (
	"green-api-test-assignment/internal/client"
	"green-api-test-assignment/internal/utils"
	envTest "green-api-test-assignment/tests"
	"testing"
)

func TestGetUrlClient(t *testing.T) {
	type args struct {
		baseUrl      string
		userInstance string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GetURLClient",
			args: args{
				baseUrl:      client.BaseURL,
				userInstance: envTest.IdInstanceMock,
			},
			want: "https://1103.api.green-api.com/waInstance" + envTest.IdInstanceMock,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetURLClient(tt.args.baseUrl, tt.args.userInstance); got != tt.want {
				t.Errorf("GetURLClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
