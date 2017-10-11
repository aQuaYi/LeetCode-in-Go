package Problem0273

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	en3 := []string{
		"thousand",
		"million",
		"billion",
	}

	res := ""
	i := -1
	for num > 0 {
		temp := lessK(num % 1000)
		if i >= 0 {
			temp += " " + en3[i]
		}

		if len(res) == 0 {
			res = temp
		} else {
			res = temp + " " + res
		}

		i++
	}

	return res
}

// 处理小于 1000 的数字
func lessK(num int) string {

	en := []string{
		"",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
		"twenty",
	}

	res := ""
	if num/100 > 0 {
		res += en[num/100] + " Hundred"
	}

	num %= 100

	if num <= 20 {
		return res + " " + en[num]
	}

	en2 := []string{
		"",
		"",
		"twenty",
		"thirty",
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
	}

	return res + " " + en2[num/10] + " " + en[num%10]
}
