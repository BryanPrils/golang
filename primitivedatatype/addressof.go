package primitivedatatype

import (
	"fmt"
)

func Address() {
	var firstName string

	firstName = "Thomas"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)
}
