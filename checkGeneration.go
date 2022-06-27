package main

import "fmt"

func main() {
	fmt.Println("Please input your age:")
	var age int
	_, err := fmt.Scanln(&age)
	if err != nil {
		fmt.Println("Incorrect input")
		return
	}

	switch {
	case age >= 1946 && age <= 1964:
		fmt.Println("Hi, boomer!")
		break
	case age >= 1965 && age <= 1980:
		fmt.Println("Hi, Mr. X!")
		break
	case age >= 1981 && age <= 1996:
		fmt.Println("Hi, millennial!")
	case age >= 1997 && age <= 2012:
		fmt.Println("Hi, zoomer!")
	case age >= 2013:
		fmt.Println("Hi, alpha!")
	case age > 2022:
		fmt.Println("You haven't been born yet")
	case age < 1946:
		fmt.Println("We can't identify your generation")
	}
}
