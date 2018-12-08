package models

import (
	"time"
)

// Project is the project model.
type Project struct {
	ID          string
	Title       string
	Description string
	ImageURI    string
	StartTime   time.Time
	State       string
}
