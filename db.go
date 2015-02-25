package empire

import (
	"database/sql"

	"github.com/remind101/empire/db"
)

// DB represents an interface for performing queries against a SQL db.
type DB interface {
	// Insert inserts a record.
	Insert(interface{}) error

	// Select performs a query and populates the interface with the
	// returned records. interface must be a pointer to a slice
	Select(interface{}, string, ...interface{}) error

	// SelectOne performs a query and populates the interface with the
	// returned record.
	SelectOne(interface{}, string, ...interface{}) error

	// Exec executes an arbitrary SQL query.
	Exec(query string, args ...interface{}) (sql.Result, error)

	// Close closes the db.
	Close() error
}

// NewDB returns a new DB instance with table mappings configured.
func NewDB(uri string) (DB, error) {
	db, err := db.NewDB(uri)
	if err != nil {
		return db, err
	}

	db.AddTableWithName(dbApp{}, "apps")

	return db, nil
}