package db

// db manager that will have function to perform operations on the database
type DBManager struct {
	db Database
}

func NewDBManager() (*DBManager, error) {
	db, err := NewDatabase("postgres", "user=postgres dbname=ctrl sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err := db.Connect(); err != nil {
		return nil, err
	}

	return &DBManager{db: db}, nil
}

func (mgr *DBManager) Close() error {
	return mgr.db.Close()
}
