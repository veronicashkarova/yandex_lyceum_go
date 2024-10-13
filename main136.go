package main

import (
	"time"
)

func NextWorkday(start time.Time) time.Time {
	var day = start.Weekday()
	var diff = 0
	switch day {
	case 5:
		diff = 3
	case 6:
		diff = 2
	default:
		diff = 1
	}

	return time.Date(start.Year(), start.Month(), start.Day()+diff, 0, 0, 0, 0, time.UTC)
}
