package main

import "fmt"

// Изменить формулу с iota, чтобы программа выводила 1 3 5 7 9 11
const (
	one = 1 + 2*iota
	three
	five
	seven
	nine
	eleven
)

func main() {
	fmt.Println(one, three, five, seven, nine, eleven)
}
