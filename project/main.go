package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"project/structs"
)

var mainView = template.Must(template.ParseFiles("templates/index.html"))
var gridView = template.Must(template.ParseFiles("templates/habit.html"))
var userData = structs.User{}


func indexHandler(w http.ResponseWriter, r *http.Request) {
	mainView.Execute(w, nil)
}

func habitViewHandler(w http.ResponseWriter, r *http.Request) {
	gridView.Execute(w, userData)
}

func loginViewHandler(w http.ResponseWriter, r *http.Request){
	params,_ := url.ParseQuery(r.URL.RawQuery)
	username := params.Get("username")
	password := params.Get("password")
	error_message := structs.CreateUser(username, password)
	w.Header().Set("Content-Type", "application/json")
	jsonData := []byte(`{"status":"OK"}`)
	if (error_message != "") {
		w.WriteHeader(http.StatusConflict)
		jsonData = []byte(`{"status":"ERROR", "error_message":` + error_message + `}`)
	} else {
		w.WriteHeader(http.StatusOK)
		userData = structs.GetUser(username)
	}
	w.Write(jsonData)
}

func updateDay(w http.ResponseWriter, r *http.Request) {
	params,_ := url.ParseQuery(r.URL.RawQuery)
	habit := params.Get("habit")
	day := params.Get("day")
	month := params.Get("month")
	level := params.Get("level")
	username := params.Get("username")
	user := structs.UpdateGrid(username, habit, day, month, level)
	fmt.Println("Updating habit: " + habit + " for user: " + username + " on day: " + day + " and month: " + month)
	userData = user
}

func createHabit(w http.ResponseWriter, r *http.Request) {
	params,_ := url.ParseQuery(r.URL.RawQuery)
	name := params.Get("name")
	goal := params.Get("goal")
	username := params.Get("username")
	keyValues := map[string]map[string]string{
		"1": {"color": params.Get("hexCode1"), "text": params.Get("level1")},
		"2": {"color": params.Get("hexCode2"), "text": params.Get("level2")},
		"3": {"color": params.Get("hexCode3"), "text": params.Get("level3")},
		"4": {"color": params.Get("hexCode4"), "text": params.Get("level4")},
	}
	habit := structs.CreateHabit(keyValues, name, goal)
	user := structs.AddHabit(username, habit)
	userData = user
	fmt.Println("Creating grid for habit: " + name + " for user: " + username + " with goal: " + goal)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/habit/", habitViewHandler)
	mux.HandleFunc("/login/", loginViewHandler)
	mux.HandleFunc("/update-day/", updateDay)
	mux.HandleFunc("/create-habit/", createHabit)
	http.ListenAndServe(":"+port, mux)
}
