package common

import (
	"github.com/vision8tech/goings/repos"
	"go.isomorphicgo.org/go/isokit"
)

// Env is used throughout the back-end side to access common components such as repositories and templates.
type Env struct {
	ProjectsRepo repos.ProjectsRepo
	TemplateSet  *isokit.TemplateSet
}
