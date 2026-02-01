package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

func ToReader(payload any) (io.Reader, error) {
	if payload == nil {
		return nil, nil
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshal payload: %w", err)
	}
	return bytes.NewReader(body), nil
}
