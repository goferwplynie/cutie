package project

import (
	"time"
)

type Project struct {
	Start      time.Time     `json:"start"`
	Deadline   time.Time     `json:"deadline"`
	Name       string        `json:"name"`
	Path       string        `json:"path"`
	Reminder   time.Duration `json:"reminder"`
	Archived   bool          `json:"archived"`
	LastActive time.Time     `json:"last_active"`
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
		Archived:   false,
	}
}
