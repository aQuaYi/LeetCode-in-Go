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
		need, pressed := 0, 0

		for i < nameSize && name[i] == c {
			need++
			i++
		}

		for j < typedSize && typed[j] == c {
			pressed++
			j++
		}

		if need > pressed {
			return false
		}

	}

	return j == typedSize
}
