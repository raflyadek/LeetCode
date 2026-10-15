func decrypt(code []int, k int) []int {
	sum := 0
	left, right := 0, k
	tempPositive := left + 1
	tempNegative := left - 1
	res := make([]int, 0)
	for left < len(code) {
		if k == 0 {
			res = append(res, 0)
			left++
		}
		if right != 0 && k > 0 {
			if tempPositive == len(code) {
				tempPositive = 0
			}
			sum += code[tempPositive]
			fmt.Printf("temp: %d\n", code[tempPositive])
			tempPositive++
			right--
			if right == 0 {
				res = append(res, sum)
				sum = 0
				right = k
				left++
				tempPositive = left + 1
			}
		}
		if right != 0 && k < 0 {
			if tempNegative < 0 {
				tempNegative = len(code) - 1
			}
			sum += code[tempNegative]
			tempNegative--
			right++
			if right == 0 {
				res = append(res, sum)
				sum = 0
				right = k
				left++
				tempNegative = left - 1
			}
		}
	}
	return res
}
