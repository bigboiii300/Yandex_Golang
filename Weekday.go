package main

import "fmt"

type Weekday int

const (
	Monday Weekday = 1 + iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func NextDay(weekday Weekday) Weekday {
	return (weekday % 7) + 1
}

func PrevDay(weekday Weekday) Weekday {
	day := (weekday % 7) - 1
	if day < 0 {
		return Saturday
	} else if day == 0 {
		return Sunday
	}
	return day
}

func main() {
	today := Sunday
	tomorrow := NextDay(today)
	yesterday := PrevDay(today)
	fmt.Println("today = ", today, "tomorrow = ", tomorrow, "yesterday = ", yesterday)
}
