package main

import (
	"html/template"
	"net/http"
	"net/url"
	"os"
	"project/structs"
)

var mainView = template.Must(template.ParseFiles("templates/index.html"))
var gridView = template.Must(template.ParseFiles("templates/habit.html"))


func indexHandler(w http.ResponseWriter, r *http.Request) {
	mainView.Execute(w, nil)
}

func habitViewHandler(w http.ResponseWriter, r *http.Request) {
	params,_ := url.ParseQuery(r.URL.RawQuery)
	username := params.Get("username")
	data := structs.GetUser(username)
	gridView.Execute(w, data)
}

func loginViewHandler(w http.ResponseWriter, r *http.Request){
	params,_ := url.ParseQuery(r.URL.RawQuery)
	username := params.Get("username")
	password := params.Get("password")
	structs.CreateUser(username, password)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData := []byte(`{"status":"OK"}`)
	w.Write(jsonData)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/habit", habitViewHandler)
	mux.HandleFunc("/login/", loginViewHandler)
	http.ListenAndServe(":"+port, mux)
}
