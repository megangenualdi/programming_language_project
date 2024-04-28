package structs

type Instance struct {
	LoggedIn User `json:"user"`
	Message string `json:"message"`
	CurrentHabit Habit `json:"currentHabit"`
}