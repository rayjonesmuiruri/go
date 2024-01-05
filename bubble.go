package main

import "fmt"

func main() {
	arr := []string{"Amber", "Idle", "Therac-25", "", "Analogy", "Antibiotic", "Characteristics"}
	fmt.Println("before>", arr)
	fmt.Println("after>", sort(arr))
}

func sort(arr []string) []string {
	swapped := false
	var count int = 0
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if len(arr[j]) < len(arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
				count++
				if !swapped{
					fmt.Println("swap not swapped")
					break
				}
			}
		}
	}
	fmt.Println("The number of checks is:", count)
	return arr
}
