package Chapter_4

import (
	"fmt"
	"runtime"
)

func main() {
	var milesDriven, gallon, gallonsUsed, miles, counter, cummulativeMilesPerGallon, milesPerGallon float32
	var input int
	flag := -1
	for input != flag {

		fmt.Print("Enter miles driven: ")
		fmt.Scan(&miles)
		counter++
		if miles < 1 {
			runtime.Breakpoint()
		}

		fmt.Print("Enter amount of gallons used: ")
		fmt.Scan(&gallon)

		milesPerGallon = miles / gallon
		fmt.Println("MilesDriven:", milesPerGallon)

		milesDriven += miles
		gallonsUsed += gallon

		cummulativeMilesPerGallon += milesPerGallon
		//fmt.Println("Cummulative Miles Per Gallon:", cummulativeMilesPerGallon)

		fmt.Println("Do you want to continue? Enter -1 to quit or any number to continue!")
		fmt.Scan(&input)

	}

	fmt.Println("Total miles driven: ", milesDriven)
	fmt.Println("Total gallons used: ", gallonsUsed)
	fmt.Println("Total number of trips made: ", counter)
	fmt.Println("Cummulative Miles Per Gallon Used: ", cummulativeMilesPerGallon)

}
