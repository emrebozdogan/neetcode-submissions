type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encodedStr := ""
	afterLenBeforeStr := "#"

	for _, str := range strs {
		encodedStr += strconv.Itoa(len(str)) + afterLenBeforeStr + str
	}

	return encodedStr
}

func (s *Solution) Decode(encoded string) []string {
	i := 0
	decodedStr := []string{}
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}

		lenOfStr, _ := strconv.Atoi(encoded[i:j])
		str := encoded[j+1:j+1+lenOfStr]
		decodedStr = append(decodedStr, str)
		i = j + 1 + lenOfStr
	}

	return decodedStr
}
