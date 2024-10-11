package handlers

import (
	"ctrl/db"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type ChapterVersionsResponse struct {
	CompanyName     string               `json:"company"`
	ProjectName     string               `json:"project"`
	ChapterName     string               `json:"chapter"`
	ChapterVersions *[]db.ChapterVersion `json:"versions"`
}

func validateID(id string) bool {
	id = strings.TrimSpace(id)
	if id == "" {
		return false
	}

	if _, err := strconv.Atoi(id); err != nil {
		return false
	}
	return true
}

func HandleChapterVersion(w http.ResponseWriter, r *http.Request) {
	// read id from path
	id := r.URL.Path[len("/chapter_versions/"):]

	if !validateID(id) {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	dbMgr, err := db.NewDBManager()
	if err != nil {
		log.Printf("Error creating DB manager: %v", err)
		http.Error(w, "Something went wrong. Please try again later.", http.StatusInternalServerError)
		return
	}
	defer dbMgr.Close()

	chaptersDetails, err := dbMgr.GetChapterDetails(id)
	if err != nil {
		log.Printf("Error getting chapter details: %v", err)
		http.Error(w, "Something went wrong. Please try again later.", http.StatusInternalServerError)
		return
	}

	if chaptersDetails == nil {
		http.Error(w, "Chapter not found", http.StatusNotFound)
		return
	}

	chapterVersions, err := dbMgr.GetChapterVersions(id)
	if err != nil {
		log.Printf("Error getting chapter versions: %v", err)
		http.Error(w, "Something went wrong. Please try again later.", http.StatusInternalServerError)
		return
	}

	response := ChapterVersionsResponse{
		CompanyName:     chaptersDetails.CompanyName,
		ProjectName:     chaptersDetails.ProjectName,
		ChapterName:     chaptersDetails.ChapterName,
		ChapterVersions: chapterVersions,
	}

	// return the response as json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
