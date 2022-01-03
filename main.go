package main

import (
	"fmt"
)

func main() {
	var sequence []int

	sequence = append(sequence, 0)
	sequence = append(sequence, 1)
	do := true
	
	for do {
		currentNumber := sequence[len(sequence)-2] + sequence[len(sequence)-1]
		sequence = append(sequence, currentNumber)

		if len(sequence) == 30 {
			do = false
		}
	}

	fmt.Println(sequence)
}
