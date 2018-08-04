package problem0478

import (
	"math"
	"math/rand"
)

// Solution object will be instantiated and called as such:
// obj := Constructor(radius, x_center, y_center);
// param_1 := obj.RandPoint();
type Solution struct {
	r, a, b, ratio float64
}

// Constructor 构建 Solution
func Constructor(radius, xCenter, yCenter float64) Solution {
	ratio := 1.
	if radius < 1 {
		radius = 1.
		ratio = 1. / radius
	}
	return Solution{
		r:     radius,
		a:     xCenter,
		b:     yCenter,
		ratio: ratio,
	}
}

// RandPoint 返回随机点
func (s *Solution) RandPoint() []float64 {
	rjd := rand.Float64()
	x := s.a + s.r*math.Sin(rjd)
	y := s.b + s.r*math.Cos(rjd)
	return []float64{x / s.ratio, y / s.ratio}
}
