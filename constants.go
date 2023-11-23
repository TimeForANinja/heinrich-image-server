package main

import (
	"os"
	"path/filepath"
)

// Configurable variables
var (
	IMAGE_DIR  = getEnv("IMAGE_DIR", "./images")
	STATIC_DIR = getEnv("STATIC_DIR", "./static")
)

// Check if an environment variable is provided, else use the default
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		path, err := filepath.Abs(value)
		if err == nil {
			return path
		}
	}
	path, err := filepath.Abs(defaultValue)
	if err != nil {
		panic(err)
	}
	return path
}
