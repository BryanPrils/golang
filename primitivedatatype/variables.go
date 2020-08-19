package primitivedatatype

func Variables() {
	var i int
	i = 42
	println(i)

	var f float32 = 3.1415
	println(f)

	firstName := "Thomas"
	println(firstName)

	c := complex(5, 4)
	println(c)

	a, b := real(c), imag(c)
	println(a, b)
}
