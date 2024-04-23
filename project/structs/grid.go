package structs

import (
	"time"
	"strconv"
	"os"
	"encoding/json"
	"fmt"
)

type Habit struct {
	Name string `json:"name"`
	Goal string `json:"goal"`
	Grids []Grid `json:"grids"`
	Key Key `json:"key"`
}


type Grid struct {
	Month string `json:"month"`
	Days []Day `json:"days"`
	Year int `json:"year"`
}


type Key struct {
	Levels []Level `json:"levels"`
}


type Level struct {
	Level int `json:"level"`
	Text string `json:"text"`
	Color string `json:"color"`
}


type Day struct {
	Number int `json:"number"`
	Completed bool `json:"completed"`
	Level Level `json:"level"`
}


func CreateHabit(keyValues map[string]map[string]string, name string, goal string) Habit {
	return Habit{
			Name: name,
			Grids: CreateGrids(),
			Key: CreateKey(keyValues),
			Goal: goal,
		}
}


func CreateKey(values map[string]map[string]string) Key {
	key := Key{
		Levels: []Level{},
	}
	for k, v := range values {
		keyInt,_ := strconv.Atoi(k)
		level := Level{
			Level: keyInt,
			Color: v["color"],
			Text: v["text"],
		}
		key.Levels = append(key.Levels, level)
	}
	return key
}


func CreateGrids() []Grid {
	currentTime := time.Now()
	currentYear := currentTime.Year()
	grids := []Grid{}
	for month := time.January; month <= time.December; month++ {
		grid := Grid{
			Month: month.String(),
			Year: currentYear,
			Days: CreateDaysByMonthYear(month, currentYear),
		}
		grids = append(grids, grid)
	}
	return grids
}


func CreateDaysByMonthYear(month time.Month, year int) []Day {
	maxDays := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
	days := []Day{}
	for num := range [32]int{} {
		num += 1
		if num > maxDays {
			break
		} else {
			day := Day{
				Number: num,
				Completed: false,
			}
			days = append(days, day)
		}
	}
	return days
}


func UpdateGrid(username string, habit string, day string, month string, level string) User {
	byteArray,err1 := os.ReadFile("/Users/mgenualdi/Desktop/projects/programming_language_project/project/jsonFiles/main.json")
	if err1 != nil {
		fmt.Println(err1)
	}

	newJson := []User{}
	err := json.Unmarshal(byteArray, &newJson)
	if err != nil {
		fmt.Println(err)
	}

	dayInt,_ := strconv.Atoi(day)
	levelInt,_ := strconv.Atoi(level)
	user := User{}

	for i := 0; i < len(newJson); i++ {
		if newJson[i].Username == username {
			user = newJson[i]
			for j := 0; j < len(user.Habits); j++ {
				if user.Habits[j].Name == habit {
					key := user.Habits[j].Key
					for k := 0; k < len(user.Habits[j].Grids); k++ {
						if user.Habits[j].Grids[k].Month == month {
							for l := 0; l < len(user.Habits[j].Grids[k].Days); l++ {
								if user.Habits[j].Grids[k].Days[l].Number == dayInt {
									user.Habits[j].Grids[k].Days[l].Completed = true
									fmt.Println(key.Levels)
									fmt.Println(levelInt)
									for m := 0; m < len(key.Levels); m++ {
										if key.Levels[m].Level == levelInt {
											fmt.Println(key.Levels[m])
											user.Habits[j].Grids[k].Days[l].Level = key.Levels[m]
										}
									}
								}
							}
						}
					}
				}
			}
			newJson[i] = user
		}
	}

	jsonStr,err2 := json.Marshal(newJson)
	if err2 != nil {
		fmt.Println(err2)
	}

	err3 := os.WriteFile("/Users/mgenualdi/Desktop/projects/programming_language_project/project/jsonFiles/main.json", jsonStr, os.ModePerm)
	if err3 != nil {
		fmt.Println(err3)
	}

	return user

}
