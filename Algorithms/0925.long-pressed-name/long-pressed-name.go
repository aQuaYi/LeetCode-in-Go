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
		count := 1
		for i+1 < nameSize && nameBytes[i] == nameBytes[i+1] {
			count++
			i++
		}

		for j < typedSize && nameBytes[i] == typedBytes[j] {
			count--
			j++
		}

		i++

		if count > 0 {
			return false
		}

	}

	return true
}
