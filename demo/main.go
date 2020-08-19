package main

import (
	"fmt"

	"example.com/demo/models"
)

func main() {
	u := models.User{
		ID:        1,
		FirstName: "Bryan",
		LastName:  "Prils",
	}
	fmt.Println(u)
}
