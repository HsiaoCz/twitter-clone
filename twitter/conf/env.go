package conf

import "os"

func GetPort(key string) string {
	port := os.Getenv(key)
	if port == "" {
		return ":8000"
	}
	return port
}
