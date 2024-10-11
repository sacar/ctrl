package handlers

import (
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

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{"chapter_id": id})
}
