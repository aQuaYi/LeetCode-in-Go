package Problem0273

import (
	"strings"
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	res := make([]string, 0, 20)
	i := 3
	k := 1000
	base := k * k * k

	for i >= 0 {
		if num/base > 0 {
			res = append(res, lessK(num/base)...)
			if i > 0 {
				res = append(res, en3[i])
			}
		}
		num %= base
		base /= k
		i--
	}

	return strings.Join(res, " ")
}

// 处理小于 1000 的数字
func lessK(num int) []string {
	res := make([]string, 0, 5)

	if num/100 > 0 {
		res = append(res, en[num/100], "Hundred")
	}

	num %= 100

	if num <= 20 {
		res = append(res, en[num])
		return res
	}

	if num/10 > 0 {
		res = append(res, en2[num/10])
	}

	num %= 10

	if num > 0 {
		res = append(res, en[num])
	}

	return res
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
	"",
	"Thousand",
	"Million",
	"Billion",
}
