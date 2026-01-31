package utils

import (
	"errors"
	utils2 "green-api-test-assignment/internal/utils"
	"reflect"
	"testing"
)

func TestReadEnv(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		args      args
		wantErr   error
		wantValue string
	}{
		{
			name: "get IDINSTANCE",
			args: args{
				key: "IDINSTANCE",
			},
			wantErr:   errors.New("переменная окружения IDINSTANCE не найдена"),
			wantValue: "",
		},
		{
			name: "get APITOKENINSTANCE",
			args: args{
				key: "APITOKENINSTANCE",
			},
			wantErr:   errors.New("переменная окружения APITOKENINSTANCE не найдена"),
			wantValue: "",
		},
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr, gotValue := utils2.ReadEnv(tt.args.key)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("ReadEnv() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
			t.Logf("value: %s", gotValue)
		})
	}
}
