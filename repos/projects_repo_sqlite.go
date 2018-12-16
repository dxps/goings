package repos

import (
	"github.com/vision8tech/goings/shared/models"
)

// NewSqliteProjectRepo is creating a new Sqlite based
// project repository instance.
func NewSqliteProjectRepo() *ProjectsRepoSqlite {

	sqliteConn := NewSqliteConnection()
	projRepo := &ProjectsRepoSqlite{}
	projRepo.Init(sqliteConn)
	return projRepo
}

// ProjectsRepoSqlite is a Sqlite based implementation of ProjectsRepo.
type ProjectsRepoSqlite struct {
	conn *RepoConnection
}

// Init is used for initializing the repo.
func (repo *ProjectsRepoSqlite) Init(conn *RepoConnection) {

	repo.conn = conn
	// table setup
	projectsTableStmt := `CREATE TABLE IF NOT EXISTS projects(
		id TEXT NOT NULL PRIMARY KEY,
		title TEXT,
		description TEXT,
		image_uri  TEXT,
		start_time TEXT,
		state TINYINT
	);`
	_, err := repo.conn.DbConnection.Exec(projectsTableStmt)
	if err != nil {
		panic(err)
	}

}

// GetProjects returns the list (slice) of existing projects.
func (repo *ProjectsRepoSqlite) GetProjects() []*models.Project {

	getAllProjectsStmt := `SELECT id, title, description, image_uri, start_time, state 
		FROM projects`
	rows, err := repo.conn.DbConnection.Query(getAllProjectsStmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// var result []*models.Project
	result := make(models.Projects, 0)
	for rows.Next() {
		item := models.Project{}
		err2 := rows.Scan(&item.ID, &item.Title, &item.Description,
			&item.ImageURI, &item.StartTime, &item.State)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, &item)
	}
	return result

}
