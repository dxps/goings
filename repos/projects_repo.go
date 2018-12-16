package repos

import (
	"github.com/vision8tech/goings/shared/models"
)

// ProjectsRepo is the contract (method set)
// for any project repository implementation.
type ProjectsRepo interface {
	Init(conn *RepoConnection)
	GetProjects() []*models.Project
}
