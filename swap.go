package piscine

func Swap(a *int, b *int) {
	x := *a
	y := *b

	value1 := y
	value2 := x

	*a = value1
	*b = value2
}
