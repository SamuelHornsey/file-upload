package main

import (
	"net/http"

	"github.com/samuelhornsey/file-upload/controllers"
)

func main () {
	// Index root
	http.Handle("/", http.FileServer(http.Dir("./public")))
	// Uploads endpoint
	http.HandleFunc("/upload", controllers.FileUpload)
	// Files endpoint
	http.HandleFunc("/files", controllers.GetFile)
	http.ListenAndServe(":8000", nil)
}