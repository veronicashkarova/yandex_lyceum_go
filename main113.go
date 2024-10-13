package main

import "fmt"

type Logger interface {
	Log(string)
}

type LogLevel string

const (
	Error LogLevel = "Error"
	Info  LogLevel = "Info"
)

type Log struct {
	Level LogLevel
}

func (log Log) Log(text string) {
	switch log.Level {
	case Error:
		fmt.Println("ERROR: ", text)
	case Info:
		fmt.Println("INFO: ", text)
	}
}
