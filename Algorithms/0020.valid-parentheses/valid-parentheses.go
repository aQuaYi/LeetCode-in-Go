package problem0020

func isValid(str string) bool {
	s := new(stack)

	for _, b := range str {
		switch b {
		case '(', '[', '{':
			s.push(b)
		case ')', ']', '}':
			if r, ok := s.pop(); !ok || r != matching[b] {
				// !ok 说明“（[{”的数量，小于")]}"的数量
				return false
			}
		}
	}

	// len(*s) > 0 说明"([{"的数量，大于")]}"的数量
	if len(*s) > 0 {
		return false
	}

	return true
}

var matching = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

type stack []rune

func (s *stack) push(b rune) {
	*s = append(*s, b)
}

func (s *stack) pop() (rune, bool) {
	if len(*s) > 0 {
		res := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return res, true
	}
	return 0, false
}
