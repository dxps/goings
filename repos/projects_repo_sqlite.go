package repos

import (
	"database/sql"
	"log"
	"time"

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
	log.Printf("NewSqliteConnection > Connected to SQLite database ('%v').\n", sqliteDB)
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
	log.Printf("ProjectsRepoSqlite.Init > %d projects exist.", len(repo.RetrieveProjects()))

}

// GetProjects returns the list (slice) of existing projects.
func (repo *ProjectsRepoSqlite) RetrieveProjects() []*models.Project {

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
		var startTime string
		err2 := rows.Scan(&item.ID, &item.Title, &item.Description,
			&item.ImageURI, &startTime, &item.State)
		if err2 != nil {
			panic(err2)
		}
		item.StartTime, _ = time.Parse(time.RFC3339, startTime)
		result = append(result, &item)
	}
	return result

}

// StoreProject stores a new project into the repository.
func (repo *ProjectsRepoSqlite) StoreProject(p *models.Project) {

	insertStmt := `INSERT INTO projects(id, title, description, image_uri, start_time, state)
	 VALUES(?,?,?,?,?,?)`
	stmt, _ := repo.conn.DbConn.Prepare(insertStmt)
	_, err := stmt.Exec(p.ID, p.Title, p.Description, p.ImageURI, p.StartTime.Format(time.RFC3339), p.State)
	if err != nil {
		log.Printf("ProjectsRepoSqlite.StoreProject > Error: '%s'\n", err)
	}

}
