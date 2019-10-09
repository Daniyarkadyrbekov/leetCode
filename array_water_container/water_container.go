package main

func maxArea(height []int) int {
	max := 0
	for i, hi:= range height{
		for j, hj := range height {
			v := calculateVolume(hi, hj, i, j)
			if max < v {
				max = v
			}
		}
	}
	return max
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

}
