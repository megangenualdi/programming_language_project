package main

import (
	"html/template"
	"net/http"
	"os"
	// "project/structs"
)

var mainView = template.Must(template.ParseFiles("templates/index.html"))
var gridView = template.Must(template.ParseFiles("templates/habit.html"))


func indexHandler(w http.ResponseWriter, r *http.Request) {
	mainView.Execute(w, nil)
}

func habitViewHandler(w http.ResponseWriter, r *http.Request) {
	gridView.Execute(w, nil)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/habit/", habitViewHandler)
	http.ListenAndServe(":"+port, mux)
	// structs.CreateUser("bgenualdi", "alsotesting")
}
