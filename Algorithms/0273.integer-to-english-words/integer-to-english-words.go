package Problem0273

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
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

		num /= 1000
		i++
	}

	return res
}

// 处理小于 1000 的数字
func lessK(num int) string {
	res := ""
	if num/100 > 0 {
		res += en[num/100] + " Hundred"
	}

	num %= 100

	if num <= 20 {
		if len(res) > 0 {
			return res + " " + en[num]
		}
		return en[num]
	}
	if len(res) > 0 {
		return res + " " + en2[num/10] + " " + en[num%10]
	}
	return en2[num/10] + " " + en[num%10]
}

var en = []string{
	"",
	"One",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
	"Eleven",
	"Twelve",
	"Thirteen",
	"Fourteen",
	"Fifteen",
	"Sixteen",
	"Seventeen",
	"Eighteen",
	"Nineteen",
	"Twenty",
}

var en2 = []string{
	"",
	"",
	"Twenty",
	"Thirty",
	"Forty",
	"Fifty",
	"Sixty",
	"Seventy",
	"Eighty",
	"Ninety",
}

var en3 = []string{
	"Thousand",
	"Million",
	"Billion",
}
