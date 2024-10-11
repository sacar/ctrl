package db

import (
	"database/sql"
	"fmt"

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

	query = fmt.Sprintf("SELECT COALESCE(json_agg(t), '[]'::json) FROM (%s) t", query)

	rows, err := psql.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result sql.NullString
	if rows.Next() {
		if err := rows.Scan(&result); err != nil {
			return nil, err
		}
	} else {
		return []interface{}{}, nil
	}

	if !result.Valid {
		return []interface{}{}, nil
	}

	return result.String, nil
}

func (psql *PSQL) Execute(query string, args ...interface{}) error {
	_, err := psql.db.Exec(query, args...)
	return err
}
