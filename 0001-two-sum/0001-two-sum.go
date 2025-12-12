func twoSum(nums []int, target int) []int {
	//target 9 so in the nums we search what index that equal 9 if we sum 
	result := make(map[int]int)

	for i, num := range nums {
		diff := target - num
		// Check if the difference already exists in the map
		if idx, found := result[diff]; found {
			return []int{i, idx}
		}
		result[num]=i
	}
	return nil
}