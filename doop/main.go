package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}

	val1, ok1 := Atoi(args[0])
	val2, ok2 := Atoi(args[2])
	operator := args[1]

	if !ok1 || !ok2 {
		return
	}

	var res int64

	switch operator {
	case "+":
		res = val1 + val2
		if (val2 > 0 && val1 > 9223372036854775807-val2) || (val2 < 0 && val1 < -9223372036854775808-val2) {
			return
		}
	case "-":
		res = val1 - val2
		if (val2 < 0 && val1 > 9223372036854775807+val2) || (val2 > 0 && val1 < -9223372036854775808+val2) {
			return
		}
	case "*":
		if val1 != 0 && val2 != 0 {
			res = val1 * val2
			if res/val1 != val2 {
				return
			}
		} else {
			res = 0
		}
	case "/":
		if val2 == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		res = val1 / val2
	case "%":
		if val2 == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		res = val1 % val2
	default:
		return
	}

	PrintInt(res)
}

func Atoi(s string) (int64, bool) {
	var res int64
	sign := int64(1)
	i := 0

	if len(s) == 0 {
		return 0, false
	}

	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	if i == len(s) {
		return 0, false
	}

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '1' {
		}
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}

		digit := int64(s[i] - '0')
		prev := res
		res = res*10 + digit

		if res < 0 || (res-digit)/10 != prev {
			if sign == -1 && res == -9223372036854775808 {
				return -9223372036854775808, true
			}
			return 0, false
		}
	}
	return res * sign, true
}

func PrintInt(n int64) {
	if n == 0 {
		os.Stdout.WriteString("0\n")
		return
	}
	if n == -9223372036854775808 {
		os.Stdout.WriteString("-9223372036854775808\n")
		return
	}

	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	res := ""
	for n > 0 {
		res = string(byte(n%10)+'0') + res
		n /= 10
	}
	os.Stdout.WriteString(sign + res + "\n")
}
