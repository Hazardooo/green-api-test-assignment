package utils

import "fmt"

func GetURLClient(baseUrl, userInstance string) string {
	return fmt.Sprintf("%s%s", baseUrl, userInstance)
}

func GetEndPointPOSTURL(baseUrl, endpoint, apiToken string) string {
	return fmt.Sprintf("%s/%s/%s", baseUrl, endpoint, apiToken)
}
