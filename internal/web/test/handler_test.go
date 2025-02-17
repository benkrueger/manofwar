// internal/web/test/handler_test.go
package test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/yourusername/hybrid-media-server/internal/web"
)

func TestNewMediaHandler(t *testing.T) {
	mediaDir := filepath.Join(os.TempDir(), "media")
	err := os.MkdirAll(mediaDir, 0755)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(mediaDir)

	handler := web.NewMediaHandler(mediaDir)

	req, err := http.NewRequest("GET", "/media/test.txt", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected status code %d, got %d", http.StatusNotFound, rr.Code)
	}
}

func TestMediaHandler(t *testing.T) {
	mediaDir := filepath.Join(os.TempDir(), "media")
	err := os.MkdirAll(mediaDir, 0755)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(mediaDir)

	// Create a test file
	testFile := filepath.Join(mediaDir, "test.txt")
	err = ioutil.WriteFile(testFile, []byte("Hello World"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	handler := web.NewMediaHandler(mediaDir)

	req, err := http.NewRequest("GET", "/media/test.txt", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	// Check the content type
	if rr.Header().Get("Content-Type") != "text/plain" {
		t.Errorf("expected content type %s, got %s", "text/plain", rr.Header().Get("Content-Type"))
	}
}
