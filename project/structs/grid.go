package structs


type Habit struct {
	Name string `json:"name"`
	Grids []Grid `json:"grids"`
}


type Grid struct {
	Month string `json:"month"`
	Days []Day `json:"days"`
	Year int `json:"year"`
	Key Key `json:"key"`
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
