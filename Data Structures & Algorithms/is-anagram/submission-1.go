func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune]int)
	tMap := make(map[rune]int)

	for _, char := range s {
		if _, exists := sMap[char]; exists {
			sMap[char] += 1
		} else {
			sMap[char] = 1
		}
	}

	for _, char := range t {
		if _, exists := tMap[char]; exists {
			tMap[char] += 1
		} else {
			tMap[char] = 1
		}
	}

	if len(sMap) != len(tMap) {
		return false
	}

	for k, valueOfsMap := range sMap {
		if valueOftMap, ok := tMap[k]; !ok || valueOfsMap != valueOftMap{
			return false
		}
	}

	return true
}
