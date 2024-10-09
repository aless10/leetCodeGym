func isValid(s string) bool {
	stack := []rune{}
	m := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for _, char := range s {
		_, ok := m[char]
		if ok {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}

			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			val, _ := m[last]
			if val != char {
				return false
			}
		}
	}

	if len(stack) > 0 {
		return false
	}
	return true
}
