package driver

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// DB :
type DB struct {
	SQL *sql.DB
	// Mgo *mgo.database
}

// DBConn :
var dbConn = &DB{}

// NewConnectionPostgres :
func NewConnectionPostgres(url string) (*DB, error) {

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	dbConn.SQL = db

	return dbConn, err
}
