package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"strings"
	"project/structs"
)

var mainView = template.Must(template.ParseFiles("templates/index.html"))
var gridView = template.Must(template.ParseFiles("templates/habit.html"))
var userData = structs.User{}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
	mainView.Execute(w, nil)
}

func habitViewHandler(w http.ResponseWriter, r *http.Request) {
	gridView.Execute(w, userData)
}


func loginViewHandler(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
		}

		username := strings.TrimSpace(r.Form.Get("username"))
    password := r.Form.Get("password")
	
		error_message := structs.CreateUser(username, password)
		if (error_message == "") {
			userData = structs.GetUser(username)
			http.Redirect(w, r, "/habit/?username="+username, http.StatusFound)
		} else {
			fmt.Println(error_message)
			mainView.Execute(w, error_message)
		}
	}
}


func updateDay(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params,_ := url.ParseQuery(r.URL.RawQuery)
	habit := params.Get("habit")
	day := params.Get("day")
	month := params.Get("month")
	level := params.Get("level")
	username := params.Get("username")
	user := structs.UpdateGrid(username, habit, day, month, level)
	fmt.Println("Updating habit: " + " for user: " + username + " on day: " + day + " and month: " + month)
	userData = user
}


func createHabit(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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
