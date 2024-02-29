package utils

import (
	"fmt"
	"net/http"
	"text/template"
)

func ParseTemplateFiles(w http.ResponseWriter, templateName string, context any, files ...string) {
	ts, er := template.ParseFiles(files...)

	if er != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		fmt.Println("Error parsing template files: ", er)
	}

	err := ts.ExecuteTemplate(w, templateName, context)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		fmt.Println("Error executing template files: ", er)
	}
}
