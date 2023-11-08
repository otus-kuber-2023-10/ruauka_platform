package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("homework.html"))

func hwHandler(w http.ResponseWriter, r *http.Request) {
	if err := tpl.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/homework.html", hwHandler)

	fmt.Println("starting server...")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
