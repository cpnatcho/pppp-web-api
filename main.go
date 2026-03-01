package main

import (
	"time"
)

func currentDate() string {
	return time.Now().Format("2006-01-02")
}
