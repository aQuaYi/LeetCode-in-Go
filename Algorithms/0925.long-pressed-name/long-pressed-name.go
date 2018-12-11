package problem0925

func isLongPressedName(name string, typed string) bool {
	if name == typed {
		return true
	}

	nameSize := len(name)
	typedSize := len(typed)

	nameBytes := []byte(name)
	typedBytes := []byte(typed)

	i, j := 0, 0
	for i < nameSize {
		c := nameBytes[i]
		pressMore := 0

		for i < nameSize && nameBytes[i] == c {
			i++
			pressMore--
		}

		for j < typedSize && typedBytes[j] == c {
			j++
			pressMore++
		}

		if pressMore < 0 {
			return false
		}

	}

	return true
}
