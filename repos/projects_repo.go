package repos

import (
	"github.com/vision8tech/goings/shared/models"
)

// ProjectsRepo is the standard contract
// for any concrete implementation of a project repository.
type ProjectsRepo interface {
	Init(conn *RepoConnection)
	GetProjects() []*models.Project
}
