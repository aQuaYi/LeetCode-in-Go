package Problem0149

import "sort"

// Point 是关于点的定义
type Point struct {
	X int
	Y int
}

func maxPoints(ps points) int {
	n := len(ps)
	if n < 3 {
		return n
	}
	max := 0

	sort.Sort(ps)

	i, j, k, temp := 0, 0, 0, 0
	for i < n-1 {
		j = i + 1
		for j < n {
			if equal(ps[i], ps[j]) {
				j++
				if i == 0 && j == n {
					return n
				}
				continue
			}
			check := checkFunc(ps[i], ps[j])

			for k = 0; k < n; k++ {
				temp = check(ps[k])
				if max < temp {
					max = temp
				}
			}
			j++
		}

		i++
	}

	return max
}

func checkFunc(p1, p2 Point) func(Point) int {
	count := 0
	return func(p Point) int {
		if (p1.X-p.X)*(p2.Y-p.Y) == (p1.Y-p.Y)*(p2.X-p.X) {
			count++
		}
		return count
	}
}

func equal(p1, p2 Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

type points []Point

func (p points) Len() int {
	return len(p)
}

func (p points) Less(i, j int) bool {
	if p[i].X == p[j].X {
		return p[i].Y < p[j].Y
	}
	return p[i].X < p[j].X
}

func (p points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
