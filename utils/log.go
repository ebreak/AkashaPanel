package utils

import (
	"fmt"
	"time"
)

func LogInfo(info string) {
	now := time.Now()
	fmt.Printf("[I] %04d-%02d-%02d %02d:%02d:%02d %s\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), info)
}
