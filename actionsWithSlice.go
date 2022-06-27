package main

import "fmt"

func main() {
	slice := make([]int, 100)
	counter := 0

	for index := range slice {
		counter++
		slice[index] = counter
	}
	// Оставляем первые и последние 10 элементов
	slice = append(slice[0:10], slice[counter-10:]...)
	fmt.Println(slice)
	// Разворачиваем слайс
	for index := 0; index < len(slice)/2; index++ {
		slice[index], slice[len(slice)-index-1] = slice[len(slice)-index-1], slice[index]
	}
	fmt.Println(slice)
}
