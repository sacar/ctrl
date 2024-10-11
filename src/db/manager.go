package db

import (
	"ctrl/config"
	"encoding/json"
	"fmt"
	"log"
)

type DBManager struct {
	db Database
}

func NewDBManager() (*DBManager, error) {
	db, err := NewDatabase("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", config.GetString("user"), config.GetString("password"), config.GetString("dbname"), config.GetString("host"), config.GetString("port")))
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
	log.Println("Getting chapter details for id:", id)
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

func (mgr *DBManager) GetChapterVersions(id string) (*[]ChapterVersion, error) {
	query := `
			SELECT
				cv.chapter_version_id,
				cv.chapter_version_number,
				cv.chapter_version_create_date,
				p.person_username,
				CASE
					WHEN cv.chapter_version_appversion = '11.0' THEN 'CC 2015'
					ELSE 'Unknown Version'
				END AS app_version
			FROM chapter_version cv
			JOIN person p ON cv.chapter_version_person_id = p.person_id
			WHERE cv.chapter_version_chapter_id = $1
			ORDER BY cv.chapter_version_number ASC
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
	var chapterVersions []ChapterVersion
	if err := json.Unmarshal([]byte(resultStr), &chapterVersions); err != nil {
		return nil, err
	}

	return &chapterVersions, nil
}
