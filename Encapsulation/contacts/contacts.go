package contacts

import "fmt"

const Email = "ikiryakov@edu.hse.ru"

var support string

func SetSupport(s string) {
	support = s
}

func GetContact() string {
	return fmt.Sprintf("%s <%s>", support, Email)
}
