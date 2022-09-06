package main

import (
	"log"
	"os"
)

func GetEnvOrExit(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("missing environment variable '%s", key)
	}
	return value
}
