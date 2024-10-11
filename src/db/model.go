package db

type Chapter struct {
	ID        int    `json:"chapter_id"`
	ProjectID int    `json:"chapter_project_id"`
	Name      string `json:"chapter_name"`
}

type ChapterDetails struct {
	ChapterName string `json:"chapter_name"`
	ProjectName string `json:"project_name"`
	CompanyName string `json:"company_name"`
}
