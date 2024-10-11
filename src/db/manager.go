package db

import (
	"encoding/json"
	"fmt"
)

// db manager that will have function to perform operations on the database
type DBManager struct {
	db Database
}

func NewDBManager() (*DBManager, error) {
	db, err := NewDatabase("postgres", "user=postgres password=sakar@123 dbname=ctrl host=localhost port=5432 sslmode=disable")
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

// function to fetch chapter from the database
func (mgr *DBManager) GetChapter(id string) (*[]Chapter, error) {
	// use json aggregate to fetch the chapter
	query := `SELECT * FROM chapter WHERE chapter_id = $1`
	result, err := mgr.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	// convert the result to string
	resultStr, ok := result.(string)
	if !ok {
		return nil, fmt.Errorf("invalid result type: %T", result)
	}

	// unmarshal the json data
	var chapters []Chapter
	if err := json.Unmarshal([]byte(resultStr), &chapters); err != nil {
		return nil, err
	}

	if len(chapters) == 0 {
		return nil, nil
	}

	return &chapters, nil
}

// function to fetch chapter versions from the database
func (mgr *DBManager) GetChapterDetails(id string) (*ChapterDetails, error) {

	query := `
			SELECT
				c.chapter_name,
				p.project_name,
				co.company_name
			FROM chapter c
			JOIN project p ON c.chapter_project_id = p.project_id
			JOIN company co ON p.project_company_id = co.company_id
			WHERE c.chapter_id = $1
			`

	result, err := mgr.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	// convert the result to string
	resultStr, ok := result.(string)
	if !ok {
		return nil, fmt.Errorf("invalid result type: %T", result)
	}

	// unmarshal the json data
	var chapterDetails []ChapterDetails
	if err := json.Unmarshal([]byte(resultStr), &chapterDetails); err != nil {
		return nil, err
	}

	if len(chapterDetails) == 0 {
		return nil, nil
	}

	return &chapterDetails[0], nil
}
