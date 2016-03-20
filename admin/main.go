package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"fmt"
)

func main() {
	// TODO: use normal router code
	
	port := "9021"

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", serveTemplate)

	log.Println("Listening...")
	http.ListenAndServe(":" + port, nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := path.Join("templates", "layout.html")
	fp := path.Join("templates", r.URL.Path)
	fmt.Println(">>>>>>", lp, fp)

	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}