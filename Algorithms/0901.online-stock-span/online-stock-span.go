package problem0901

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
	s.prices = append(s.prices, price)
	res := 1
	for i := len(s.prices) - 2; 0 <= i && s.prices[i] <= price; i-- {
		res++
	}
	return res
}
