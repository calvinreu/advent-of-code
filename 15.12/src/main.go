package main

import (
	"fmt"
	"strconv"
)

var startingNumbers []float64
var numbers map[float64]uint = make(map[float64]uint)

func main() {

	var input string = ""

	fmt.Println("Enter starting values to end setting values press e")

	for input != "e" {
		fmt.Scanln(&input)
		number, err := strconv.ParseFloat(input, 64)
		if err == nil {
			startingNumbers = append(startingNumbers, number)
		}
	}

	input = ""

	fmt.Println(startingNumbers)

	for _, num := range startingNumbers {
		numbers[num] = 0
	}

	fmt.Println("Enter number to exit press e")

	for i := 1; input != "e"; i++ {
		fmt.Scanln(&input)
		number, err := strconv.ParseFloat(input, 64)
		if err == nil {
			if numbers[number] == 0 {
				fmt.Println("0")
			} else {
				fmt.Println((uint)(i)-numbers[number], "Truns ago")
			}
			numbers[number] = (uint)(i)
		}
	}

}
