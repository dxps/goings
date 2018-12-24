package repos

import (
	"github.com/vision8tech/goings/shared/models"
)

// ProjectsRepo is the contract (method set)
// for any project repository implementation.
type ProjectsRepo interface {

	// Init is used for the repo initialization.
	Init(conn *RepoConnection)

	// RetrieveProjects retrieves all the existing projects from the repo.
	RetrieveProjects() ([]*models.Project, error)

	// RetrieveProjectByID retrieves an existing project by its identifier.
	RetrieveProjectByID(id string) (*models.Project, error)

	// StoreProject persists a new project into the repo.
	StoreProject(p *models.Project) error

	// UpdateProjectByID is responsible with updating an existing project
	// found in the repo based on its identifier.
	UpdateProjectByID(id string, p *models.Project) error

	// DeleteProjectByID is deleting an existing project based on its identifier
	DeleteProjectByID(id string) error

	// Uninit is used from a clean/graceful shutdown.
	Uninit()
}
