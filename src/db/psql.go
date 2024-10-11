package db

import "database/sql"

type PSQL struct {
	db *sql.DB
}

func (psql *PSQL) Connect() error {
	var err error
	psql.db, err = sql.Open("postgres", "user=postgres dbname=ctrl sslmode=disable")
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
