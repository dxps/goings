package common

import "github.com/vision8tech/goings/repos"

// Env is used throughout the project to access common components such as repositories.
type Env struct {
	ProjectsRepo repos.ProjectsRepo
}
