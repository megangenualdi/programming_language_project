package structs

import (
	"time"
	"strconv"
)

type Habit struct {
	Name string `json:"name"`
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
	Color string `json:"color"`
}


type Day struct {
	Number int `json:"number"`
	Completed bool `json:"completed"`
	Level Level `json:"level"`
}


func CreateHabit(keyValues map[string]string, name string) Habit {
	if (len(keyValues) == 0) {
		keyValues = map[string]string{
			"1": "#A3E4D7",
			"2": "#48C9B0",
			"3": "#17A589",
			"4": "#0E6251",
		}
	}
	return Habit{
			Name: name,
			Grids: CreateGrids(),
			Key: CreateKey(keyValues),
		}
}


func CreateKey(values map[string]string) Key {
	key := Key{
		Levels: []Level{},
	}
	for k, v := range values {
		keyInt,_ := strconv.Atoi(k)
		level := Level{
			Level: keyInt,
			Color: v,
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
		if num == maxDays {
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


