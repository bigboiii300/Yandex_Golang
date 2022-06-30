package main

import "fmt"

func RemoveDuplicates(input []string) []string {
	output := make([]string, len(input))
	result := make([]string, len(input))
	copy(output, input)
	var index int

	inputSet := make(map[string]int, len(input))
	for _, value := range input {
		if _, ok := inputSet[value]; !ok {
			inputSet[value]++
			result[index] = value
			index++
		}
	}
	return result
}

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}
	fmt.Println(input)

	fmt.Println(RemoveDuplicates(input))
}
