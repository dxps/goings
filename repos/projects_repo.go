package repos

import (
	"github.com/vision8tech/goings/shared/models"
)

// ProjectsRepo is the contract (method set)
// for any project repository implementation.
type ProjectsRepo interface {
	Init(conn *RepoConnection)

	RetrieveProjects() ([]*models.Project, error)

	RetrieveProjectByID(id string) (*models.Project, error)

	StoreProject(p *models.Project)

	Uninit()
}
