package repos

import (
	"database/sql"
	"log"

	"github.com/vision8tech/goings/shared/models"

	// register the SQLite driver
	_ "github.com/mattn/go-sqlite3"
)

const sqliteDB = "./goings.sqlitedb"

// locally reused for clean shutdown (uninit)
var sqliteConn *RepoConnection

// NewSqliteConnection creates a connection to SQLite database.
func NewSqliteConnection() *RepoConnection {

	conn, err := sql.Open("sqlite3", sqliteDB)
	if err != nil {
		log.Fatal(err)
	}
	repoConn := RepoConnection{DbConn: conn}
	log.Printf("(repos) NewSqliteConnection > Connected to SQLite database ('%v').\n", sqliteDB)
	sqliteConn = &repoConn
	return &repoConn

}

// UninitSqliteConnection should be called in a graceful shutdown case.
func UninitSqliteConnection() {
	sqliteConn.DbConn.Close()
	log.Println("SQLite database connection closed.")
}

// ProjectsRepoSqlite is a Sqlite based implementation of ProjectsRepo.
type ProjectsRepoSqlite struct {
	conn *RepoConnection
}

// NewSqliteProjectRepo is creating an SQLite based project repository instance.
func NewSqliteProjectRepo() *ProjectsRepoSqlite {

	sqliteConn := NewSqliteConnection()
	projRepo := &ProjectsRepoSqlite{}
	projRepo.Init(sqliteConn)
	return projRepo

}

// Init is used for initializing the repo.
func (repo *ProjectsRepoSqlite) Init(conn *RepoConnection) {

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
	_, err := repo.conn.DbConn.Exec(projectsTableStmt)
	if err != nil {
		panic(err)
	}
	log.Printf("[ProjectsRepoSqlite] Init(conn) > %d projects exist.", len(repo.GetProjects()))

}

// GetProjects returns the list (slice) of existing projects.
func (repo *ProjectsRepoSqlite) GetProjects() []*models.Project {

	getAllProjectsStmt := `SELECT id, title, description, image_uri, start_time, state 
		FROM projects`
	rows, err := repo.conn.DbConn.Query(getAllProjectsStmt)
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

// StoreProject stores a new project into the repository.
func (repo *ProjectsRepoSqlite) StoreProject(project *models.Project) {

	insertStmt := `INSERT INTO projects(id, title, description, image_uri, start_time, state)
	 VALUES($1, $2, $3, $4, $5, $6)`
	repo.conn.DbConn.QueryRow(insertStmt)
}
