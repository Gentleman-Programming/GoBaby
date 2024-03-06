package utils

import (
	"net/http"
	"text/template"
)

// HTTPMethod represents an enumeration of HTTP methods
type HTTPMethod int

const (
	// GET HTTP method
	GET HTTPMethod = iota
	// POST HTTP method
	POST
	// PUT HTTP method
	PUT
	// Other methods can be added similarly
)

// String returns the string representation of the HTTPMethod
func (m HTTPMethod) String() string {
	switch m {
	case GET:
		return http.MethodGet
	case POST:
		return http.MethodPost
	case PUT:
		return http.MethodPut
	default:
		return ""
	}
}

func IsValidHTTPMethod(method string, acceptedMethod string, w http.ResponseWriter) bool {
	if method == acceptedMethod {
		return true
	}

	w.Header().Set("Allow", acceptedMethod)

	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

	return false
}

var acceptedMethods = []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}

func CheckIfPath(w http.ResponseWriter, r *http.Request, path string) {
	if r.URL.Path != path {
		http.NotFound(w, r)
		return
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
