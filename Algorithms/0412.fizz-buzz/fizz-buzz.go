package problem0412

import "strconv"

func fizzBuzz(n int) []string {
	res := make([]string, n)

	for i := range res {
		x := i + 1
		switch {
		case x%15 == 0:
			res[i] = "FizzBuzz"
		case x%5 == 0:
			res[i] = "Buzz"
		case x%3 == 0:
			res[i] = "Fizz"
		default:
			res[i] = strconv.Itoa(x)
		}
	}

	return res
}
