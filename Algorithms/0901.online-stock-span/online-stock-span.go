package problem0901

// StockSpanner object will be instantiated and called as such:
// obj := Constructor();
// param_1 := obj.Next(price);
type StockSpanner struct {
	prices []int
	days   []int
	today  int
}

// Constructor is
func Constructor() StockSpanner {
	return StockSpanner{}
}

// Next is
func (s *StockSpanner) Next(p int) int {
	s.today++
	i := len(s.prices) - 1
	for ; i >= 0; i-- {
		if s.prices[i] > p {
			break
		}
	}
	i++
	s.prices = append(s.prices[:i], p)
	s.days = append(s.days[:i], s.today)
	if i == 0 {
		return s.today
	}
	return s.today - s.days[i-1]
}
