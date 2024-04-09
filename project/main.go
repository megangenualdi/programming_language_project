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

func updateDay(w http.ResponseWriter, r *http.Request) {
	params,_ := url.ParseQuery(r.URL.RawQuery)
	habit := params.Get("habit")
	day := params.Get("day")
	month := params.Get("month")
	level := params.Get("level")
	username := params.Get("username")
	structs.UpdateGrid(username, habit, day, month, level)
	fmt.Println("Updating habit: " + habit + " for user: " + username + " on day: " + day + " and month: " + month)

}

func createHabit(w http.ResponseWriter, r *http.Request) {
// 	params,_ := url.ParseQuery(r.URL.RawQuery)
// 	habit := params.Get("habit")
// 	username := params.Get("username")
// 	structs.CreateHabit(habit, username)
// 	fmt.Println("Creating grid for habit: " + habit + " for user: " + username)
// }

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
	// mux.HandleFunc("/create-habit/", createHabit)
	http.ListenAndServe(":"+port, mux)
}
