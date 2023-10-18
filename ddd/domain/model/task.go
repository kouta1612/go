package model

import "time"

type Task struct {
	id          uint64
	title       string
	completed   bool
	createdAt   time.Time
	completedAt *time.Time
}

func Create(title string) *Task {
	return &Task{
		id:          1, // TODO
		title:       title,
		completed:   false,
		createdAt:   time.Now(),
		completedAt: nil,
	}
}

func (task *Task) Complete() {
	now := time.Now()
	task.completed = true
	task.completedAt = &now
}
