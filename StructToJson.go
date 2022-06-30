package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type PersonJ struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"Дата рождения"`
}

func main() {
	marshal, err := json.Marshal(PersonJ{
		Name:  "Alex",
		Email: "alex@yandex.ru",
	})
	if err != nil {
		return
	}
	fmt.Printf("Man %v\n", string(marshal))
}
