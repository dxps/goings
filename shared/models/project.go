package models

import (
	"time"
)

// Project is the model of a project.
type Project struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageURI    string    `json:"image_uri"`
	StartTime   time.Time `json:"start_time"`
	State       string    `json:"state"`
}

// Projects is a slice of 0-N references to Project objects.
type Projects []*Project
