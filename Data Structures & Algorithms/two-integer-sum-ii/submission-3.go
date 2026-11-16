func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for {
		result := numbers[left] + numbers[right]
		if result > target {
			right--
		} else if result < target {
			left++
		} else {
			return []int{left+1, right+1}
		}
	}
	return []int{}
}
