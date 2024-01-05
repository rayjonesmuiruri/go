package main

import "fmt"

func main() {
	//check if a year is a leap year or not
	var year int
	fmt.Println("Enter a year to check if it is leap year: ")
	fmt.Scan(&year)
	var rem4, rem100, rem400 int

	rem4 = year % 4
	rem100 = year % 100
	rem400 = year % 400

	if (rem4 == 0 && rem100 != 0) || rem400 == 0 {
		fmt.Printf("%d is a leap year!\n", year)
	} else {
		fmt.Print("Not a leap year\n")
	}

}
