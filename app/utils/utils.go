package utils

import (
	"log"
	"os"
)

func CheckError(err error) {
    if err != nil {
        log.Fatalf("Error: %v", err)
    }
}

func GetEnv(key string) string {
    return os.Getenv(key)
}
