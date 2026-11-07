func twoSum(nums []int, target int) []int {
  leftovers := map[int]int{}

  for i := 0; i < len(nums); i++ {
	difference := target - nums[i]
	if _, exists := leftovers[difference]; exists {
		return []int{leftovers[difference], i}
	}
	leftovers[nums[i]] = i
  }
  return []int{}
}