package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	power := 1
	result := 0

	for {
		digit, end := getDigit(x, power)
		if end {
			break
		}
		result = result * 10 + digit
		power++
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		result = 0
	}
	return result
}

func tenInPower(pow int) int {
	res := 1
	for i := 0; i < pow; i++{
		res *= 10
	}

	return res
}

func getDigit(number int, power int) (int, bool) {
	number = number / tenInPower(power - 1)
	if number == 0 {
		return 0, true
	}
	number = number % 10
 	return number, false
}

func main() {

	x := -1534236

	fmt.Println(reverse(x))
}