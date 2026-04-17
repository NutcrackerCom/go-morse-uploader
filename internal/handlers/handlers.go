package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	html, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "error in reading html", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(html)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "error parsing form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "error in getting file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "error reading file", http.StatusBadRequest)
		return
	}

	convertedData, err := service.ConvertText(string(data))
	if err != nil {
		http.Error(w, "error converting file", http.StatusInternalServerError)
		return
	}

	root, err := os.OpenRoot("uploads")
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	defer root.Close()

	ext := filepath.Ext(handler.Filename)
	name := time.Now().UTC().Format("20060102-150405.000000000") + ext
	dst, err := root.Create(name)
	if err != nil {
		http.Error(w, "error in creating file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	str := fmt.Sprintf("%s %s", time.Now().UTC().String(), convertedData)
	if _, err := dst.WriteString(str); err != nil {
		http.Error(w, "error in writing to file", http.StatusInternalServerError)
		return
	}

}
