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


func indexHandler(w http.ResponseWriter, r *http.Request) {
	mainView.Execute(w, instance)
	instance.Message = ""
}


func habitViewHandler(w http.ResponseWriter, r *http.Request) {
	gridView.Execute(w, instance)
	instance.Message = ""
}


func createAccountHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
		}

		username := strings.TrimSpace(r.Form.Get("username"))
    password := strings.TrimSpace(r.Form.Get("password"))

		if (username == "" || password == "") {
			fmt.Println("The username and password fields cannot be empty.")
			instance.Message = "The username and password fields cannot be empty."
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		
	
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
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
		}

		username := strings.TrimSpace(r.Form.Get("username"))
    password := strings.TrimSpace(r.Form.Get("password"))

		if (username == "" || password == "") {
			fmt.Println("The username and password fields cannot be empty.")
			instance.Message = "The username and password fields cannot be empty."
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		
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


func logoutViewHandler(w http.ResponseWriter, r *http.Request) {
	instance = structs.Instance{}
	http.Redirect(w, r, "/", http.StatusFound)
}


func updateDay(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating day")
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
		}
		habit := strings.TrimSpace(r.Form.Get("habit"))
		day := strings.TrimSpace(r.Form.Get("day"))
		month := strings.TrimSpace(r.Form.Get("month"))
		level := strings.TrimSpace(r.Form.Get("level"))
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
	fmt.Println("Creating habit")
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
		}
		name := strings.TrimSpace(r.Form.Get("name"))
		goal := strings.TrimSpace(r.Form.Get("goal"))
		username := instance.LoggedIn.Username
	  keyValues := map[string]map[string]string{
			"1": {"color": strings.TrimSpace(r.Form.Get("hexCode1")), "text": parseFormValue(r.Form.Get("level1"), "1")},
			"2": {"color": strings.TrimSpace(r.Form.Get("hexCode2")), "text": parseFormValue(r.Form.Get("level2"), "2")},
			"3": {"color": strings.TrimSpace(r.Form.Get("hexCode3")), "text": parseFormValue(r.Form.Get("level3"), "3")},
			"4": {"color": strings.TrimSpace(r.Form.Get("hexCode4")), "text": parseFormValue(r.Form.Get("level4"), "4")},
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


func parseFormValue(value string, defaultValue string) string {
	if (strings.TrimSpace(value) == "") {
		return defaultValue
	} else {
		return strings.TrimSpace(value)
	}
}


// PORT=<enter port> FILEPATH=<enter path to json file>  go run main.go
func main() {
	port := os.Getenv("PORT")

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/habit/", habitViewHandler)
	mux.HandleFunc("/create-account/", createAccountHandler)
	mux.HandleFunc("/login/", loginViewHandler)
	mux.HandleFunc("/logout/", logoutViewHandler)
	mux.HandleFunc("/update-day/", updateDay)
	mux.HandleFunc("/create-habit/", createHabit)
	http.ListenAndServe(":"+port, mux)
}
