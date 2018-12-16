package models

import (
	"time"
)

// Project is the model of a project.
type Project struct {
	ID          string
	Title       string
	Description string
	ImageURI    string
	StartTime   time.Time
	State       string
}

// Projects is a slice of 0-N references to Project objects.
type Projects []*Project
