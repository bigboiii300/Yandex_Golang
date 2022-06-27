package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name        string
	Age         int
	lastVisited time.Time
}

func updateLastVisited(person *Person) {
	person.lastVisited = time.Now()
}

func main() {
	person := Person{
		Name:        "John",
		Age:         21,
		lastVisited: time.Time{},
	}
	fmt.Println(person.Name, person.Age, person.lastVisited.Format("January 2, 2006"))
	updateLastVisited(&person)
	fmt.Println(person.Name, person.Age, person.lastVisited.Format("January 2, 2006"))
}
