package repos

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// RepoConnection is providing a database connection setup:
// the connection itself, the state.
type RepoConnection struct {
	DbConnection *sql.DB
}

const sqliteDB = "./goings.sqlitedb"

// locally reused for clean shutdown (uninit)
var sqliteConn *RepoConnection

// NewSqliteConnection creates a connection to SQLite database.
func NewSqliteConnection() *RepoConnection {

	conn, err := sql.Open("sqlite3", sqliteDB)
	if err != nil {
		log.Fatal(err)
	}
	repoConn := RepoConnection{DbConnection: conn}
	log.Printf("Connected to SQLite database ('%v').\n", sqliteDB)
	sqliteConn = &repoConn
	return &repoConn

}

// UninitSqliteConnection should be called in a graceful shutdown case.
func UninitSqliteConnection() {
	sqliteConn.DbConnection.Close()
	log.Println("SQLite database connection closed.")
}
