package main

import (
	"log"
	"sort"
)

func checkAvailable() {
	lc := newLeetCode()

	problems := lc.Problems

	list := make([]int, 0, len(problems))

	for _, p := range problems {
		if !p.IsAccepted && !isAvailable(p) {
			list = append(list, p.ID)
		}
	}

	sort.Ints(list)
	new := &unavailable{List: list}

	showUnavailableChange(new)

	new.save()
}

func showUnavailableChange(new *unavailable) {
	old := readUnavailable()

	size := max(new.List[len(new.List)-1], old.List[len(old.List)-1])

	for i := 0; i <= size; i++ {
		newHas := new.has(i)
		oldHas := old.has(i)

		if newHas && !oldHas {
			log.Printf("新发现：第 %d 题，无法使用 Go 解答", i)
		}

		if !newHas && oldHas {
			log.Printf("好消息：第 %d 题，可以使用 Go 解答了", i)
		}
	}
}

func isAvailable(p problem) bool {
	res := true

	// 若本题无法使用 Go 语言解答，getFunction 会 panic
	// 此 defer 就是用于处理此种情况
	defer func() {
		if err := recover(); err != nil {
			res = false
		}
	}()

	getFunction(p.link())

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
