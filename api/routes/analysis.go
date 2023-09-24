package routes

import (
	"net/http"
)

func Analysis(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
	default:
		w.WriteHeader(404)
	}
}
