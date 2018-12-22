package repos

import (
	"database/sql"
)

// RepoConnection provides the repository connection details:
// the connection itself, the state.
type RepoConnection struct {

	// database connection reference
	DbConn *sql.DB
}
