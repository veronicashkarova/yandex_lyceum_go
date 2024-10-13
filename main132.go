package main

import (
	"time"
)

func FormatTimeToString(timestamp time.Time, format string) string {
	result := timestamp.Format(format)
	return result
}