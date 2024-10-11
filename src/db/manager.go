package db

// db manager that will have function to perform operations on the database
type DBManager struct {
	db Database
}

func NewDBManager(driver string) (*DBManager, error) {
	db, err := NewDatabase(driver)
	if err != nil {
		return nil, err
	}
	return &DBManager{db: db}, nil
}
