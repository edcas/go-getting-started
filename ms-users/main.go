package main

import (
	"fmt"

	"github.com/company/ms-users/models"
)

func main() {
	user := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}

	fmt.Println(user)
}
