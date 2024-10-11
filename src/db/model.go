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

type ChapterVersion struct {
	ChapterVersionID int    `json:"chapter_version_id"`
	Version          int    `json:"chapter_version_number"`
	CreatedAt        string `json:"chapter_version_create_date"`
	CreatedBy        string `json:"person_username"`
	Description      string `json:"chapter_version_description"`
}
