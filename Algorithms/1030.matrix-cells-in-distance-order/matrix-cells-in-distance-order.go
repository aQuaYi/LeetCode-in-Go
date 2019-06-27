package problem1030

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}

var tmp map[int][][]int = make(map[int][][]int)
var result [10000][]int

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {

	for indexR := 0; indexR < R; indexR++ {
		for indexC := 0; indexC < C; indexC++ {
			index := abs(indexR-r0) + abs(indexC-c0)
			tmp[index] = append(tmp[index], []int{indexR, indexC})
		}
	}
	indexResult := 0
	for index := 0; index <= R+C; index++ {
		val := tmp[index]
		num := len(val)
		end := indexResult + num
		copy(result[indexResult:end], val)
		indexResult = end

		delete(tmp, index)
	}

	return result[:indexResult]
}
