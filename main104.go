package main

import "time"

type Task struct {
	summary (string)
	description (string)
	deadline (time.Time)
	priority (int)
}

func (p Task) IsOverdue () bool {
	return time.Now().Compare(p.deadline)<0
}

func (p Task) IsTopPriority () bool {
	return p.priority > 3
}