package problem1032

// StreamChecker check letters
type StreamChecker struct {
	dic  [][]*char
	time int
}

type char struct {
	index, queryTime int // alphabet order
	isEnd            bool
	next             *char
}

func makeChars(word string) *char {
	n := len(word)
	head := &char{
		index: int(word[0] - 'a'),
	}
	end := head
	for i := 1; i < n; i++ {
		end.next = &char{
			index: int(word[i] - 'a'),
		}
		end = end.next
	}
	end.isEnd = true
	end.next = head
	return head
}

// Constructor returns StreamChecker
func Constructor(words []string) StreamChecker {
	dic := make([][]*char, 26)
	for _, word := range words {
		l := makeChars(word)
		dic[l.index] = append(dic[l.index], l)
	}
	return StreamChecker{
		dic: dic,
	}
}

// Query returns true if letter in words
func (sc *StreamChecker) Query(letter byte) bool {
	time := sc.time
	sc.time++
	index := int(letter - 'a')

	chars := sc.dic[index]
	sc.dic[index] = make([]*char, 0, len(chars))

	res := false
	for _, c := range chars {
		c.queryTime = time
		res = res || (c.isEnd && c.check())
		c = c.next
		sc.dic[c.index] = append(sc.dic[c.index], c)
	}

	return res
}

func (c *char) check() bool {
	head := c.next
	qt := head.queryTime
	for !head.isEnd {
		head = head.next
		qt++
		if qt != head.queryTime {
			return false
		}
	}
	return true
}
