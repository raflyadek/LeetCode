func isPalindrome(x int) bool {
	stringInt := strconv.Itoa(x)
	halfLenArray := len(stringInt)/2
	lastIndexArray := len(stringInt)
	for i := 0; i < halfLenArray; i++ {
		if stringInt[i] != stringInt[lastIndexArray-i-1] {
			return false
		}
	}
	return true
}