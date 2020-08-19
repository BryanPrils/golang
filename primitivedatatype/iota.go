package primitivedatatype

const (
	first = iota
	second
)

const (
	third  = iota
	fourth = iota + 6
)

func Iota() {
	println(first, second, third, fourth)
}
