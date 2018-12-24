package sqlite

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/vision8tech/goings/repos"
	"github.com/vision8tech/goings/shared/models"

	// register the SQLite driver
	_ "github.com/mattn/go-sqlite3"
)

// SQLite database file path
const sqliteFile = "./goings.sqlitedb"

//
// newSqliteRepoConnection is used internally to create the connection to SQLite database.
//
func newSqliteRepoConnection() *repos.RepoConnection {

	sqliteDb, err := sql.Open("sqlite3", sqliteFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("NewSqliteRepoConnection > Connected to database ('%v').\n", sqliteFile)
	repoConn := repos.RepoConnection{Db: sqliteDb}
	return &repoConn

}

//
// ProjectsRepoSqlite is a Sqlite based implementation of ProjectsRepo.
//
type ProjectsRepoSqlite struct {
	conn *repos.RepoConnection
}

//
// NewSqliteProjectRepo is creating an SQLite based project repository instance.
//
func NewSqliteProjectRepo() *ProjectsRepoSqlite {

	sqliteConn := newSqliteRepoConnection()
	projRepo := &ProjectsRepoSqlite{}
	projRepo.Init(sqliteConn)
	return projRepo

}

//
// Init is used for initializing the repo.
//
func (repo *ProjectsRepoSqlite) Init(conn *repos.RepoConnection) {

	repo.conn = conn
	// create table, if not already exists
	projectsTableStmt := `CREATE TABLE IF NOT EXISTS projects(
		id TEXT NOT NULL PRIMARY KEY,
		title TEXT,
		description TEXT,
		image_uri  TEXT,
		start_time TEXT,
		state TINYINT
	);`
	_, err := repo.conn.Db.Exec(projectsTableStmt)
	if err != nil {
		log.Printf("ProjectsRepoSqlite.Init > Error creating the projects table (if not already exists): %s\n", err)
		panic(err)
	}
	projects, err := repo.RetrieveProjects()
	if err != nil {
		log.Printf("ProjectsRepoSqlite.Init > Error counting the projects: %s\n", err)
	}
	log.Printf("ProjectsRepoSqlite.Init > %d projects exist.", len(projects))

}

//
// Uninit is used during the graceful shutdown of the system.
// It closes the database connection for now.
//
func (repo *ProjectsRepoSqlite) Uninit() {

	_ = repo.conn.Db.Close()
	log.Println("ProjectsRepoSqlite.Uninit > SQLite database connection closed.")

}

//
// RetrieveProjects returns the list (slice) of existing projects.
//
func (repo *ProjectsRepoSqlite) RetrieveProjects() ([]*models.Project, error) {

	getAllProjectsStmt := `SELECT id, title, description, image_uri, start_time, state FROM projects`
	rows, err := repo.conn.Db.Query(getAllProjectsStmt)
	if err != nil {
		log.Printf("ProjectsRepoSqlite.RetrieveProjects > Error on query: %s\n", err)
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	result := make(models.Projects, 0)
	for rows.Next() {
		item := models.Project{}
		var startTime string
		err = rows.Scan(&item.ID, &item.Title, &item.Description, &item.ImageURI, &startTime, &item.State)
		if err != nil {
			return nil, err
		}
		item.StartTime, _ = time.Parse(time.RFC3339, startTime)
		result = append(result, &item)
	}
	return result, nil

}

//
// RetrieveProjectByID returns a project by id.
//
func (repo *ProjectsRepoSqlite) RetrieveProjectByID(id string) (*models.Project, error) {

	var proj = models.Project{}
	rows, err := repo.conn.Db.Query(`SELECT id, title, description, image_uri, start_time, state
		FROM projects WHERE id=?`, id)
	defer func() { _ = rows.Close() }()
	if err != nil {
		log.Printf("ProjectsRepoSqlite.RetrieveProjectByID > Error retrieving data: %s\n", err)
		return nil, err
	}
	if rows.Next() {
		var startTime string
		err = rows.Scan(&proj.ID, &proj.Title, &proj.Description, &proj.ImageURI, &startTime, &proj.State)
		if err != nil {
			log.Printf("ProjectsRepoSqlite.RetrieveProjectByID > Error scanning the row: %s\n", err)
			return &proj, err
		}
		proj.StartTime, _ = time.Parse(time.RFC3339, startTime)
		return &proj, nil
	}
	return nil, nil

}

//
// StoreProject stores a new project into the repository.
//
func (repo *ProjectsRepoSqlite) StoreProject(p *models.Project) error {

	if p == nil {
		return errors.New("provided project reference is nil")
	}
	insertStmt := `INSERT INTO projects(id, title, description, image_uri, start_time, state)
	 VALUES(?,?,?,?,?,?)`
	stmt, _ := repo.conn.Db.Prepare(insertStmt)
	_, err := stmt.Exec(p.ID, p.Title, p.Description, p.ImageURI, p.StartTime.Format(time.RFC3339), p.State)
	if err != nil {
		log.Printf("ProjectsRepoSqlite.StoreProject > Error: '%s'\n", err)
	}
	return err

}

//
// UpdateProjectByID updates an existing project.
//
func (repo *ProjectsRepoSqlite) UpdateProjectByID(id string, p *models.Project) error {

	if p == nil {
		return errors.New("provided project reference is nil")
	}
	foundProject, err := repo.RetrieveProjectByID(id)
	if foundProject == nil {
		return errors.New("non-existing project")
	} else if err != nil {
		return err
	}
	// finally, now the project exists, let's update it
	updateStmt := `UPDATE projects SET title=?, description=?, image_uri=?, state=? WHERE id=?`
	stmt, _ := repo.conn.Db.Prepare(updateStmt)
	_, err = stmt.Exec(p.Title, p.Description, p.ImageURI, p.State, id)
	if err != nil {
		log.Printf("ProjectsRepoSqlite.UpdateProjectByID > Error: '%s'\n", err)
		return err
	} else {
		return nil // all good, no errors
	}

}

//
// DeleteProjectByID updates an existing project.
//
func (repo *ProjectsRepoSqlite) DeleteProjectByID(id string) error {

	foundProject, err := repo.RetrieveProjectByID(id)
	if foundProject == nil {
		return errors.New("non-existing project")
	} else if err != nil {
		return err
	}
	// finally, now the project exists, let's delete it
	deleteStmt := `DELETE FROM projects WHERE id=?`
	stmt, _ := repo.conn.Db.Prepare(deleteStmt)
	_, err = stmt.Exec(id)
	if err != nil {
		log.Printf("ProjectsRepoSqlite.DeleteProjectByID > Error: '%s'\n", err)
		return err
	} else {
		return nil // all good, no errors
	}

}
