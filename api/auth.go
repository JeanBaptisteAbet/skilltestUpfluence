package main

import "net/http"

func Protect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//here we will implement the auth service
		next.ServeHTTP(w, r)
		return

	})
}
