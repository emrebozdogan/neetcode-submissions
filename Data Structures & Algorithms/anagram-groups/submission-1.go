func groupAnagrams(strs []string) [][]string {
	groupedByID := make(map[[26]int][]string)

	for _, str := range strs {
		var id [26]int
		for _, c := range str {
			id[c - 'a']++
		}
		groupedByID[id] = append(groupedByID[id], str)
	}

	groupedAnagrams := make([][]string, 0, len(groupedByID))

	for _, group := range groupedByID {
		groupedAnagrams = append(groupedAnagrams, group)
	}
	return groupedAnagrams
}
