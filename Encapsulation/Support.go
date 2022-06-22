package main

import (
	"awesomeProject1/contacts"
	"fmt"
)

func main() {
	contacts.SetSupport("Служба поддержки")
	fmt.Println(contacts.GetContact())
	fmt.Println("Адрес электронной почты:", contacts.Email)
}
