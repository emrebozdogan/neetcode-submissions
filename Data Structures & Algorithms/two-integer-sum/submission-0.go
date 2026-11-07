func twoSum(nums []int, target int) []int {
  leftovers := map[int]int{}
  var difference int
  var index int

  for i := 0; i < len(nums); i++ {
	index = i
	difference = target - nums[i]
	if _, exists := leftovers[difference]; !exists {
		leftovers[nums[i]] = i
	} else {
		break
		}
  }
	return []int{leftovers[difference], index}
}