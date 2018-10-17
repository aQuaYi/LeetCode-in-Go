package problem0901

import (
	"sort"
)

// StockSpanner object will be instantiated and called as such:
// obj := Constructor();
// param_1 := obj.Next(price);
type StockSpanner struct {
	prices []int
}

// Constructor is
func Constructor() StockSpanner {
	ps := make([]int, 0, 10000)
	return StockSpanner{
		prices: ps,
	}
}

// Next is
func (s *StockSpanner) Next(price int) int {
	i := sort.SearchInts(s.prices, price)
	res := i + 1
	s.prices = append(s.prices, price)
	copy(s.prices[i+1:], s.prices[i:])
	s.prices[i] = price
	return res
}
