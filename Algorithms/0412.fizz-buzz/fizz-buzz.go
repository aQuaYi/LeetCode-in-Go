package Problem0412

import "strconv"

func fizzBuzz(n int) []string {
	res := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		temp := ""

		if i%3 == 0 {
			temp = "Fizz"
		}

		if i%5 == 0 {
			temp += "Buzz"
		}

		if temp == "" {
			temp = strconv.Itoa(i)
		}

		res = append(res, temp)
	}

	return res
}
