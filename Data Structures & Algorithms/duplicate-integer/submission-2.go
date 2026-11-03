func hasDuplicate(nums []int) bool {
	duplicates := map[int]bool{};
	for _, value := range nums {
		if _, exists := duplicates[value]; exists {
			return true;
		}
		duplicates[value] = true;
	} 
	return false;
}
