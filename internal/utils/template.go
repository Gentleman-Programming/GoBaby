package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func ParseTemplateFiles(w http.ResponseWriter, templateName string, context any, funcToTemplate template.FuncMap, files ...string) {
	ts, er := template.ParseFiles(files...)
	if er != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		fmt.Println("Error parsing template files: ", er)
		return
	}

	if funcToTemplate != nil {
		ts = ts.Funcs(funcToTemplate)
	}

	err := ts.ExecuteTemplate(w, templateName, context)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		fmt.Println("Error executing template files: ", err)
	}
}
