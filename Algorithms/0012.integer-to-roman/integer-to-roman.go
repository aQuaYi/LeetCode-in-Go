package problem0012

func intToRoman(num int) string {

	d := [4][]string{
		[]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		[]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		[]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		[]string{"", "M", "MM", "MMM"},
	}
	return d[3][num/1000] +
		d[2][num/100%10] +
		d[1][num/10%10] +
		d[0][num%10]
}
