package project

import (
	"time"
)

type Project struct {
	Start      time.Time
	Deadline   time.Time
	Name       string
	Path       string
	Reminder   time.Duration
	Archived   bool
	LastActive time.Time
}

func New(deadline time.Time, name, path string, reminder time.Duration) *Project {
	now := time.Now()
	return &Project{
		Start:      now,
		Deadline:   deadline,
		Name:       name,
		Path:       path,
		Reminder:   reminder,
		LastActive: now,
	}
}
