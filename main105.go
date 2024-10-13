package main

import "time"

type Task struct {
	summary     (string)
	description (string)
	deadline    (time.Time)
	priority    (int)
}

type Note struct {
	title (string)
	text  (string)
}

type ToDoList struct {
	name  (string)
	tasks ([]Task)
	notes ([]Note)
}

func (p ToDoList) TasksCount() int {
	return len(p.tasks)
}

func (p ToDoList) NotesCount() int {
	return len(p.notes)
}

func (p Task) IsTopPriority() bool {
	return p.priority > 3
}

func (p ToDoList) CountTopPrioritiesTask() int {
	var x int = 0
	for _, task := range p.tasks {
		if task.IsTopPriority() {
			x++
		}
	}
	return x
}

func (p Task) IsOverdue() bool {
	return time.Now().Compare(p.deadline) > 0
}

func (p ToDoList) CountOverdueTasks() int {
	var y int = 0
	for _, task := range p.tasks {
		if task.IsOverdue() {
			y++
		}
	}
	return y
}
