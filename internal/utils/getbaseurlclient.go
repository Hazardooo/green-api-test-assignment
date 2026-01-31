package utils

import "fmt"

func GetURLClient(baseUrl, userInstance string) string {
	return fmt.Sprintf("%s%s", baseUrl, userInstance)
}

func GetEndPointURL(baseUrl, userInstance, ApiToken, endpoint string) string {
	return fmt.Sprintf("%s%s/%s/%s", baseUrl, userInstance, endpoint, ApiToken)
}
