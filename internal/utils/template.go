package utils

import (
	"net/http"
	"text/template"
)

func ParseTemplateFiles(w http.ResponseWriter, files ...string) {
	ts, er := template.ParseFiles(files...)

	if er != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	err := ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
