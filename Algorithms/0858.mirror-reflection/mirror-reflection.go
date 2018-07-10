package problem0858

// 在坐标系中，画上间距为 p 的网格
// 按照题目中的镜像关系标注网格中的点为 0,1,2
// 可以总结出以下规律
func mirrorReflection(p int, q int) int {
	l := lcm(p, q)
	if (l/q)%2 == 0 {
		return 2
	}
	return (l / p) % 2
}

// p 和 q 的最小公倍数
// p >= q
func lcm(p, q int) int {
	return p * q / gcd(p, q)
}

// p 和 q 的最大公约数
// p >= q
func gcd(p, q int) int {
	for q != 0 {
		p, q = q, p%q
	}
	return p
}
