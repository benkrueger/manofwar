package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/benkrueger/manofwar/internal/web"
)

func main() {
	// Define and parse command-line flags
	port := flag.String("port", "8080", "server port")
	mediaDir := flag.String("mediadir", "./media", "directory containing media files")
	flag.Parse()

	// Use environment variables if they are set
	if envPort := os.Getenv("SERVER_PORT"); envPort != "" {
		*port = envPort
	}
	if envMediaDir := os.Getenv("MEDIA_DIR"); envMediaDir != "" {
		*mediaDir = envMediaDir
	}

	// Validate the media directory
	absMediaDir, err := filepath.Abs(*mediaDir)
	if err != nil {
		log.Fatalf("Error getting absolute path for media directory: %v", err)
	}

	http.HandleFunc("/media/", web.NewMediaHandler(absMediaDir))

	log.Printf("Server started at http://localhost:%s", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
