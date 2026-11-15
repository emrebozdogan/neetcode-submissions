func isPalindrome(s string) bool {
	var sb strings.Builder
	
	sb.Grow(len(s))

	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9'){
			sb.WriteByte(c)
		} else if c >= 'A' && c <= 'Z' {
			sb.WriteByte(c + ('a' - 'A')) 
		}
	}
	newS := sb.String()
	length := len(newS)
	if length <= 1 {
		return true
	}
	j := length-1
	for i := 0; i <= (length / 2); i++ {
		if newS[i] != newS[j] {
			return false
		}
		j--
	}

	return true
}
