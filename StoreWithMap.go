package main

import "fmt"

func main() {
	priceList := map[string]int{"хлеб": 50, "молоко": 100, "масло": 200, "колбаса": 500, "соль": 20, "огурцы": 200, "сыр": 600, "ветчина": 700, "буженина": 900, "помидоры": 250, "рыба": 300, "хамон": 1500}
	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	var totalPrice int
	for k, v := range priceList {
		if v > 500 {
			fmt.Println(k)
		}
	}

	for _, value := range order {
		fmt.Println(value, " стоит ", priceList[value])
		totalPrice += priceList[value]
	}
	fmt.Println(totalPrice)
}
