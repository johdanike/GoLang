package piscine

func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b
	remainder := a % b
	*mod = remainder
}
