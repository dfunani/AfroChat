package utils

import (
	"fmt"
	"os"
)

func GetEnv(key string) string {
	if value := os.Getenv(key); value != "" {
		fmt.Println("Environment variable:", key, "=", value)
		return value
	}
	panic(fmt.Sprintf("Environment variable %s is not set", key))
}
