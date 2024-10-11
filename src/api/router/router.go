package router

import (
	"ctrl/api/handlers"
	"net/http"
)

func SetupRouter() {
	// setup routes
	http.HandleFunc("/chapter_versions/{id}", handlers.HandleChapterVersion)

	// start server
	http.ListenAndServe(":8080", nil)
}
