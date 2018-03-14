package problem0735

func asteroidCollision(asteroids []int) []int {
	res := make([]int, 0, len(asteroids))
	stack := make([]int, 0, len(asteroids))

	for _, a := range asteroids {
		if a > 0 {
			stack = append(stack, a)
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] <= -a {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if last == -a {
				a = 0
				break
			}
		}

		if len(stack) == 0 && a != 0 {
			res = append(res, a)
			continue
		}
	}

	res = append(res, stack...)

	return res
}
