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
	user.Habits = []Habit{
		CreateHabit(map[string]string{}, "Reading"),
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