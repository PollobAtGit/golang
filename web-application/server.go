package main

import (
	"html/template"
	"net/http"
	"os"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(rw http.ResponseWriter, r *http.Request) {
	tpl.Execute(rw, nil)
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", indexHandler)

	http.ListenAndServe(":"+port, mux)
}
