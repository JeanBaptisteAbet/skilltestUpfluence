package routes

import (
	"net/http"
	"skillsTestUpfluence/api/controllers"
)

func Analysis(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		controllers.Analysis(w, r)
	default:
		w.WriteHeader(404)
	}
}
