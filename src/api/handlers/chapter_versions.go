package handlers

import (
	"ctrl/db"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dbMgr.Close()

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{"chapter_id": id})
}
