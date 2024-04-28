package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"project/structs"
)


var mainView = template.Must(template.ParseFiles("templates/index.html"))
var gridView = template.Must(template.ParseFiles("templates/habit.html"))
var instance = structs.Instance{}


func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
	mainView.Execute(w, instance)
	instance.Message = ""
}


func habitViewHandler(w http.ResponseWriter, r *http.Request) {
	gridView.Execute(w, instance)
	instance.Message = ""
}


func createAccountHandler(w http.ResponseWriter, r *http.Request){
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
			fmt.Println("Account successfully created.")
			instance.Message = "Account successfully created. Please login."
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			fmt.Println(error_message)
			instance.Message = error_message
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
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
		
		findUser := structs.GetUser(username)
		if (findUser.Username != "") {
			if (findUser.Password == password) {
				instance.LoggedIn = findUser
				instance.CurrentHabit = instance.LoggedIn.Habits[0]
				http.Redirect(w, r, "/habit/", http.StatusFound)
			} else {
				fmt.Println("Incorrect password")
				instance.Message = "Incorrect password"
				http.Redirect(w, r, "/", http.StatusFound)
			}
		} else {
			fmt.Println("This username does not exist.")
			instance.Message = "This username does not exist."
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}


func updateDay(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Println("Updating day")
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
		}
		habit := r.Form.Get("habit")
		day := r.Form.Get("day")
		month := r.Form.Get("month")
		level := r.Form.Get("level")
		if (level == "") {
			return
		}
		username := instance.LoggedIn.Username
		fmt.Println(habit, day, month, level, username)
		user, currentHabit := structs.UpdateGrid(username, habit, day, month, level)
		fmt.Println("Updating habit: " + " for user: " + username + " on day: " + day + " and month: " + month)
		instance.CurrentHabit = currentHabit
		instance.LoggedIn = user
		http.Redirect(w, r, "/habit/", http.StatusFound)
	}
}


func createHabit(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Println("Creating habit")
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
		}
		name := r.Form.Get("name")
		goal := r.Form.Get("goal")
		username := instance.LoggedIn.Username
	  keyValues := map[string]map[string]string{
			"1": {"color": r.Form.Get("hexCode1"), "text": r.Form.Get("level1")},
			"2": {"color": r.Form.Get("hexCode2"), "text": r.Form.Get("level2")},
			"3": {"color": r.Form.Get("hexCode3"), "text": r.Form.Get("level3")},
			"4": {"color": r.Form.Get("hexCode4"), "text": r.Form.Get("level4")},
		}
		habit := structs.CreateHabit(keyValues, name, goal)
		user := structs.AddHabit(username, habit)
		instance.LoggedIn = user
		instance.CurrentHabit = habit
		instance.Message = "Successfully created grid"
		http.Redirect(w, r, "/habit/", http.StatusFound)
		fmt.Println("Creating grid for habit: " + name + " for user: " + username + " with goal: " + goal)
	}
}


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/habit/", habitViewHandler)
	mux.HandleFunc("/create-account/", createAccountHandler)
	mux.HandleFunc("/login/", loginViewHandler)
	mux.HandleFunc("/update-day/", updateDay)
	mux.HandleFunc("/create-habit/", createHabit)
	http.ListenAndServe(":"+port, mux)
}
