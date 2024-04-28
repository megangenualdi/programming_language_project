package structs

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Habits []Habit `json:"habits"`
}

func CreateUser(username string, password string) string {
	user := User{
		Username: username,
		Password: password,
	}
	keyValues := map[string]map[string]string{
		"1": {"color": "#A3E4D7", "text": "Read less than ten minutes"},
		"2": {"color": "#48C9B0", "text": "Read more than ten minutes"},
		"3": {"color": "#17A589", "text": "Read 30 minutes or less"},
		"4": {"color": "#0E6251", "text": "Read more than 30 minutes"},
	}
	user.Habits = []Habit{
		CreateHabit(keyValues, "Reading", "Read at least 30 minutes per day"),
	}

	byteArray,err1 := os.ReadFile("/Users/megangenualdi/programming_language_project/project/jsonFiles/main.json")
	if err1 != nil {
		fmt.Println(err1)
	}
	newJson := []User{}
	err := json.Unmarshal(byteArray, &newJson)
	if err != nil {
		fmt.Println(err)
	}

	error_message := ""
	newUser := true
	for i := 0; i < len(newJson); i++ {
		if (newJson[i].Username == username) {
			newUser = false
			fmt.Println("The entered username is taken")
			error_message = "The entered username is taken"
		}
	}
	if error_message != "" {
		return error_message
	}

	if (newUser && error_message == "") {
		newJson = append(newJson, user)
		jsonStr,err2 := json.Marshal(newJson)
		if err2 != nil {
			fmt.Println(err2)
		}

		err3 := os.WriteFile("/Users/megangenualdi/programming_language_project/project/jsonFiles/main.json", jsonStr, os.ModePerm)
		if err3 != nil {
			fmt.Println(err3)
	}	
	}
	return error_message
}


func AddHabit(username string, habit Habit) User {
	byteArray,err1 := os.ReadFile("/Users/megangenualdi/programming_language_project/project/jsonFiles/main.json")
	if err1 != nil {
		fmt.Println(err1)
	}
	newJson := []User{}
	err := json.Unmarshal(byteArray, &newJson)
	if err != nil {
		fmt.Println(err)
	}

	user := User{}
	for i := 0; i < len(newJson); i++ {
		if newJson[i].Username == username {
			user = newJson[i]
			user.Habits = append(user.Habits, habit)
			newJson[i] = user
		}
	}
	user.Habits = append(user.Habits, habit)
	fmt.Println("Adding habit: " + habit.Name + " for user: " + username)

	jsonStr,err2 := json.Marshal(newJson)
	if err2 != nil {
		fmt.Println(err2)
	}

	err3 := os.WriteFile("/Users/megangenualdi/programming_language_project/project/jsonFiles/main.json", jsonStr, os.ModePerm)
	if err3 != nil {
		fmt.Println(err3)
	}

	return user

}

func GetUser(username string) User {
	byteArray,err1 := os.ReadFile("/Users/megangenualdi/programming_language_project/project/jsonFiles/main.json")
	if err1 != nil {
		fmt.Println(err1)
	}
	newJson := []User{}
	err := json.Unmarshal(byteArray, &newJson)
	if err != nil {
		fmt.Println(err)
	}
	
	user := User{}
	for i := 0; i < len(newJson); i++ {
		if newJson[i].Username == username {
			fmt.Println("User found")
			user = newJson[i]
		}
	}
	return user
}
