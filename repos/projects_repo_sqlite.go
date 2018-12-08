package repos

import (
	"github.com/vision8tech/goings/shared/models"
)

// ProjectsRepoSqlite is a Sqlite based implementation of ProjectsRepo.
type ProjectsRepoSqlite struct {
}

// GetProjects() returns the list (slice) of existing projects.
func (repo ProjectsRepoSqlite) GetProjects() []*models.Project {

	// todo
	return nil

}
