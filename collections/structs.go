package collections

import "fmt"

func Structs() {

	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	u := user{
		ID:        1,
		FirstName: "Bryan",
		LastName:  "Prils",
	}
	fmt.Println(u)
}
