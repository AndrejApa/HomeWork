package main

import "fmt"

func main() {
	arr := []uint{
		28, 33, 16,
		7, 5, 88,
	}

	max := arr[0]
	for _, value := range arr {
		if value < max {
			max = value // found another smaller value, replace previous value in max
		}
	}

	fmt.Println("The biggest/largest value is : ", max)
}
