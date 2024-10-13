package main

import (
	"strconv"
	"time"
)

func TimeAgo(pastTime time.Time) string {
	var end = time.Now()
	var diff = end.Sub(pastTime).Hours()
	var str = strconv.FormatFloat(diff, 'f', 0, 64)
	return str + " hours ago"
}
