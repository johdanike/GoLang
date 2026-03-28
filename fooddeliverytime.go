package piscine

type food struct {
	preptime int
}

func setTime(fd *food, time int) {
	fd.preptime = time
}

func FoodDeliveryTime(order string) int {
	time := &food{}

	if order == "burger" {
		setTime(time, 15)
	} else if order == "chips" {
		setTime(time, 10)
	} else if order == "nuggets" {
		setTime(time, 12)
	} else {
		return 404
	}

	return time.preptime
}
