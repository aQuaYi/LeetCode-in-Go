package problem0478

import (
	"fmt"
	"math"
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

// RandPoint 返回随机点
func (s *Solution) RandPoint() []float64 {
	len := math.Sqrt(rand.Float64()) * s.r
	deg := rand.NormFloat64() * 2 * math.Pi
	x := s.a + len*math.Cos(deg)
	y := s.b + len*math.Sin(deg)
	fmt.Printf("len = %f, x = %f, y = %f\n", len, x, y)
	return []float64{x, y}
}
