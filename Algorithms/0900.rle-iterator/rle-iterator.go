package problem0900

// RLEIterator object will be instantiated and called as such:
// obj := Constructor(A);
// param_1 := obj.Next(n);
type RLEIterator struct {
	es []entry
}

type entry struct {
	count, number int
}

// Constructor is
func Constructor(a []int) RLEIterator {
	size := len(a)
	es := make([]entry, 0, size/2)
	for i := 0; i+1 < size; i += 2 {
		if a[i] == 0 {
			continue
		}
		es = append(es, entry{
			count:  a[i],
			number: a[i+1],
		})
	}
	return RLEIterator{
		es: es,
	}
}

// Next is
func (r *RLEIterator) Next(n int) int {
	var prev entry
	for len(r.es) > 0 && r.es[0].count <= n {
		n -= r.es[0].count
		prev, r.es = r.es[0], r.es[1:]
	}

	if len(r.es) == 0 && n > 0 {
		return -1
	}

	if n == 0 {
		return prev.number
	}

	r.es[0].count -= n
	return r.es[0].number
}
