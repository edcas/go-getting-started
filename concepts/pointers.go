package pointers

import "fmt"

func pointers() {
	var name *string = new(string)
	*name = "John"
	fmt.Println(*name)

	firstName := "Arthur"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Tricia"
	fmt.Println(ptr, *ptr)
}
