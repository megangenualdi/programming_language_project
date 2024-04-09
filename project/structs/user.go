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

func CreateUser(username string, password string) {
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

	byteArray,err1 := os.ReadFile("/Users/mgenualdi/Desktop/projects/programming_language_project/project/jsonFiles/main.json")
	if err1 != nil {
		fmt.Println(err1)
	}
	newJson := []User{}
	err := json.Unmarshal(byteArray, &newJson)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(newJson); i++ {
		if newJson[i].Username == username {
			fmt.Println("Username already exists")
			return
		}
	}

	newJson = append(newJson, user)
	jsonStr,err2 := json.Marshal(newJson)
	if err2 != nil {
		fmt.Println(err2)
	}

	err3 := os.WriteFile("/Users/mgenualdi/Desktop/projects/programming_language_project/project/jsonFiles/main.json", jsonStr, os.ModePerm)
	if err3 != nil {
		fmt.Println(err3)
	}
}

func GetUser(username string) User {
	byteArray,err1 := os.ReadFile("/Users/mgenualdi/Desktop/projects/programming_language_project/project/jsonFiles/main.json")
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
