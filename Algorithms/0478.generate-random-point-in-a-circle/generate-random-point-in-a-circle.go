package problem0478

import (
	"math/rand"
)

// Solution object will be instantiated and called as such:
// obj := Constructor(radius, x_center, y_center);
// param_1 := obj.RandPoint();
type Solution struct {
	r, a, b float64
}

// Constructor 构建 Solution
func Constructor(radius, xCenter, yCenter float64) Solution {
	return Solution{
		r: radius,
		a: xCenter,
		b: yCenter,
	}
}

// RandPoint 返回圆内的随机点
func (s *Solution) RandPoint() []float64 {
	xFactor, yFactor := 1., 1.
	for xFactor*xFactor+yFactor*yFactor > 1 {
		xFactor = 2*rand.Float64() - 1
		yFactor = 2*rand.Float64() - 1
	}
	x := s.a + s.r*xFactor
	y := s.b + s.r*yFactor
	return []float64{x, y}
}
