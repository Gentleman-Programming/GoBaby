package utils

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

func ParseTemplateFiles(w http.ResponseWriter, templateName string, context any, funcToTemplate template.FuncMap, embedFS embed.FS, files ...string) {
	ts, er := template.ParseFS(embedFS, files...)
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

func TransformToTemplateFuncMap(key string, f interface{}) template.FuncMap {
	return template.FuncMap{
		key: f,
	}
}

var (
	EmptyFuncMap = template.FuncMap{}
	EmptyStruct  = struct{}{}
)
