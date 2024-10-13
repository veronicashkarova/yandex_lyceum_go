package main

import (
    "fmt"
    "time"
)

func TimeDifference(start, end time.Time) time.Duration {
    start := time.Date()
    end := time.Date()
    diff := end.Sub(start)
    fmt.Println(diff)
}