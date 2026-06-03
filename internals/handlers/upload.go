package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

const uploadDir = "./uploads"
const maxUploadSize = 5 << 20 // 5 MB

var allowedExtensions = map[string]bool{
	".jpg": true, ".jpeg": true, ".png": true, ".gif": true,
	".webp": true, ".svg": true,
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		http.Error(w, "file too large or invalid form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "missing file field", http.StatusBadRequest)
		return
	}
	defer file.Close()

	ext := filepath.Ext(header.Filename)
	if !allowedExtensions[ext] {
		http.Error(w, fmt.Sprintf("file type %s not allowed", ext), http.StatusBadRequest)
		return
	}

	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		http.Error(w, "unable to create upload directory", http.StatusInternalServerError)
		return
	}

	filename := uuid.New().String() + ext
	dst, err := os.Create(filepath.Join(uploadDir, filename))
	if err != nil {
		http.Error(w, "unable to save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "unable to write file", http.StatusInternalServerError)
		return
	}

	fileURL := fmt.Sprintf("/uploads/%s", filename)

	// Clean up old files after 24 hours (in production use a proper cleanup job)
	go func() {
		time.Sleep(24 * time.Hour)
		os.Remove(filepath.Join(uploadDir, filename))
	}()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"url": fileURL})
}
