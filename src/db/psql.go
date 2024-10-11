package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PSQL struct {
	driver  string
	connStr string
	db      *sql.DB
}

func (psql *PSQL) Connect() error {
	var err error
	psql.db, err = sql.Open(psql.driver, psql.connStr)
	if err != nil {
		return err
	}
	return nil
}

func (psql *PSQL) Close() error {
	return psql.db.Close()
}

func (psql *PSQL) Query(query string, args ...interface{}) (interface{}, error) {
	// TODO: implement
	return nil, nil
}

func (psql *PSQL) Execute(query string, args ...interface{}) error {
	// TODO: implement
	return nil
}
