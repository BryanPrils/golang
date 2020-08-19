package primitivedatatype

import (
	"fmt"
)

func Pointers() {
	var firstName *string = new(string)

	*firstName = "Thomas"
	fmt.Println(*firstName)
}
