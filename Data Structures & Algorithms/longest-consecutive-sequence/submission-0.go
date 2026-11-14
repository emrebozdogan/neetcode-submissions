func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool)

	for _, num := range nums {
		numMap[num] = true
	}
	longest := 0
	for num, _ := range numMap {
		if numMap[num - 1] {
			continue
		}
		length := 1
		for numMap[num + length] {
			length++
		}
		if longest < length {
			longest = length
		}
	}

	return longest
}
