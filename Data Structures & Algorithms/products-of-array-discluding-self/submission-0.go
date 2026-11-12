func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	prefix := 1

	for i, v := range nums {
		output[i] = prefix
		prefix *= v
	}

	suffix := 1

	for j := len(nums)-1; j >= 0; j-- {
		output[j] *= suffix
		suffix *= nums[j]
	}

	return output
}
