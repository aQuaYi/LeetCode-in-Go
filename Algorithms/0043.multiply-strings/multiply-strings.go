package Problem0043

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	base := []byte(num1)
	count := []byte(num2)
	temp := make([]int, len(num1)+len(num2))

	for i := 0; i < len(base); i++ {
		for j := 0; j < len(count); j++ {
			temp[i+j+1] += int(base[i]-'0') * int(count[j]-'0')
		}
	}

	for i := len(temp) - 1; i > 0; i-- {
		temp[i-1] += temp[i] / 10
		temp[i] = temp[i] % 10
	}

	if temp[0] == 0 {
		temp = temp[1:]
	}

	res := make([]byte, len(temp))
	for i := 0; i < len(temp); i++ {
		res[i] = '0' + byte(temp[i])
	}
	return string(res)
}
