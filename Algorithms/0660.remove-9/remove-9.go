package Problem0660

func newInteger(n int) int {
	i := 0
	for n > 0 {
		i = next(i)
		n--
	}
	return i
}

func next(n int) int {
	inc := 1

}
