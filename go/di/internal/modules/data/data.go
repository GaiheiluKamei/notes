package data

import (
	// import the mysql driver
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"errors"
)

const (
	// default person id (return on error)
	defaultPersonID = 0
)

var (
	db *sql.DB

	// ErrNotFound is returned when the no records where matched by the query
	ErrNotFound = errors.New("not found")
)