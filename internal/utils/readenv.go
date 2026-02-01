package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env файла не найдено")
	}
}

func ReadEnv(key string) (error, string) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fmt.Errorf("переменная окружения %s не найдена", key), ""
	}
	return nil, value
}
