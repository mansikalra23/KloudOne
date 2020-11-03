package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := Name{"Mansi", "Kalra"}
	template, _ := template.ParseFiles("hello.gohtml")
	template.Execute(w, name) // for error
}

type Name struct {
	fName, lName string
}
