func hasDuplicate(nums []int) bool {
	duplicates := make(map[int]struct{});
	for _, value := range nums {
		if _, exists := duplicates[value]; exists {
			return true;
		}
		duplicates[value] = struct{}{};
	} 
	return false;
}
