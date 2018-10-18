package problem0901

// StockSpanner object will be instantiated and called as such:
// obj := Constructor();
// param_1 := obj.Next(price);
type StockSpanner struct {
	prices []int // stack
	days   []int // stack
}

// Constructor is
func Constructor() StockSpanner {
	ps := make([]int, 1, 10000)
	ds := make([]int, 1, 10000)
	ps[0] = 1<<63 - 1 // more than max-price
	ds[0] = -1        // the day before first-day
	return StockSpanner{
		prices: ps,
		days:   ds,
	}
}

// Next is
func (s *StockSpanner) Next(price int) int {
	i := len(s.prices) - 1
	// s.prices[i] 中保存了昨天的 price
	// s.days[i] 中保存了昨天的 日期
	today := s.days[i] + 1

	for i >= 0 {
		if s.prices[i] > price {
			// 深入 s.prices 栈，直到找到比 price 大的价格
			break
		}
		i--
	}

	res := today - s.days[i]
	i++
	s.prices = append(s.prices[:i], price)
	s.days = append(s.days[:i], today)
	return res
}
