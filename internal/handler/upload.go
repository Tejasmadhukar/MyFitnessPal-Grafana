package handler

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(60 << 20)

	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No file or bad file was sent"))
		return
	}
	defer file.Close()

	fileName := strings.Split(header.Filename, ".")

	if fileName[1] != "csv" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Server only accepts a csv file"))
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error reading file " + err.Error()))
		return
	}

	err = os.WriteFile("internal/assets/data/"+header.Filename, data, 0644)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not save file"))
		return
	}

}
