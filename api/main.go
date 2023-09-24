package main

import (
	"fmt"
	"net/http"
	"skillsTestUpfluence/api/routes"
)

func main() {
	Startserver()
}

func Startserver() {

	http.Handle("/analysis", Protect(http.HandlerFunc(routes.Analysis)))

	//--------------- start server and print error ---------------//
	fmt.Println(http.ListenAndServe(":8080", nil)) //TODO cors handler

}
