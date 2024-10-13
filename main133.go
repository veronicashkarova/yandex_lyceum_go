package main

import "time"

func ParseStringToTime(dateString string, format string) (time.Time, error) {
	return time.Parse(format,dateString)
}