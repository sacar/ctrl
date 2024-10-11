package db

import "fmt"

// interface for database operations
type Database interface {
	Connect() error
	Close() error
	Query(query string, args ...interface{}) (interface{}, error)
	Execute(query string, args ...interface{}) error
}

func NewDatabase(driver string) (Database, error) {
	switch driver {
	case "psql":
		return &PSQL{}, nil
	default:
		return nil, fmt.Errorf("invalid database driver: %s", driver)
	}
}
