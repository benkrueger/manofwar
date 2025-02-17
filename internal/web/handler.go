package web

import (
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// NewMediaHandler creates a new handler for serving files from the specified directory
func NewMediaHandler(mediaDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the filename from the URL path
		filename := strings.TrimPrefix(r.URL.Path, "/media/")
		filepath := filepath.Join(mediaDir, filename)

		// Log the request
		log.Printf("Request for file: %s", filepath)

		// Open the requested file
		file, err := os.Open(filepath)
		if err != nil {
			// File not found, return 404
			http.NotFound(w, r)
			return
		}
		defer file.Close()

		// Determine the content type based on the file extension
		extension := filepath.Ext(filename)
		mimeType := mime.TypeByExtension(extension)
		if mimeType == "" {
			mimeType = "application/octet-stream"
		}

		// Set the content type header
		w.Header().Set("Content-Type", mimeType)

		// Serve the file content
		http.ServeContent(w, r, filename, file.ModTime(), file)
	}
}
