package problem0925

func isLongPressedName(name, typed string) bool {
	if name == typed {
		return true
	}

	nameSize := len(name)
	typedSize := len(typed)

	i, j := 0, 0
	for i < nameSize {
		c := name[i]
		pressMore := 0

		for i < nameSize && name[i] == c {
			i++
			pressMore--
		}

		for j < typedSize && typed[j] == c {
			j++
			pressMore++
		}

		if pressMore < 0 {
			return false
		}

	}

	return j == typedSize
}
