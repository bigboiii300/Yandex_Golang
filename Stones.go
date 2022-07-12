package main

import "fmt"

func main() {
	var count int
	_, err := fmt.Scan(&count)
	if err != nil {
		return
	}
	var input, max, counter int
	for i := 0; i < count; i++ {
		_, err := fmt.Scan(&input)
		if err != nil {
			return
		}
		if input == 1 {
			counter++
		}
		if max <= counter {
			max = counter
		}
		if input == 0 {
			counter = 0
		}
	}
	fmt.Println(max)
}
