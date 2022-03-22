package main

import (
	"fmt"
	"strconv"
)

func main() {
	var maxDigit int
	fmt.Scanf("%d", &maxDigit)
	result := getPossibleNumbers(maxDigit)
	fmt.Print(result)
}

func getPossibleNumbers(maxDigit int) []int {
	result := make([]int, 0)
	for i := 1000; i < 9999; i++ {
		strNumber := strconv.Itoa(i)

		sum := 0
		for _, value := range []byte(strNumber) {
			value, err := strconv.Atoi(string(value))
			if err != nil {
				continue
			}

			if value > maxDigit {
				continue
			}

			sum += value
		}

		if sum != 21 {
			continue
		}

		result = append(result, i)
	}

	return result
}
