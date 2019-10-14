package main

import "fmt"

var max = 0

func maxArea(height []int) int {
	getValume(height)
	return max
}

func getValume(height []int) {
	if len(height) < 1 {
		return
	}
	localVolume := calculateVolume(height[0], height[len(height) - 1], 0, len(height) - 1)
	if localVolume > max {
		max = localVolume
	}
	if height[0] > height[len(height) - 1] {
		getValume(height[:len(height) - 1])
	}else{
		getValume(height[1:])
	}
}

func calculateVolume(hi, hj, i, j int) int {
	h := 0
	if hi > hj {
		h = hj
	}else{
		h = hi
	}
	w := i - j
	if w < 0 {
		w *= -1
	}
	return w * h
}

func main() {
	fmt.Println(maxArea([]int{1,1}))
}

