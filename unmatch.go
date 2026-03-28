package piscine

// func Unmatch(a []int) int {

// 	for index := 0; index < len(a); index++ {

// 		counter := 0

// 		for count := 0; count < len(a); count++ {
// 			if a[index] == a[count] {
// 				counter++
// 			}
// 		}

// 		if counter%2 != 0 {
// 			return a[index]
// 		}

// 	}

// 	return -1

// }

func Unmatch(a []int) int {
	for index := 0; index < len(a); index++ {

		counter := 0

		for count := 0; count < len(a); count++ {
			if a[index] == a[count] {
				counter++
			}
		}

		// if counter%2 != 0 {
		// 	return a[index]
		// }

		if counter != 2 {
			return a[index]
		}

	}

	return -1
}
