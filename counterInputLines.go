package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(cnt *int) {
	*cnt += 1
}

func main() {
	// Получаем читателя пользовательского ввода
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Interaction counter")

	count := 0
	for {
		fmt.Print("-> ")
		// Считываем введённую пользователем строку. Программа ждёт, пока пользователь введёт строку
		_, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		countLines(&count)

		fmt.Printf("User input %d lines\n", count)
	}
}
