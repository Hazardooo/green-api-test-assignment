package utils

import "fmt"

func GetUrlClient(baseUrl, userInstance string) string {
	return fmt.Sprintf("%s%s", baseUrl, userInstance)
}
