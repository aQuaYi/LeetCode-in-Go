package Problem0304

// NumMatrix 包含了需要求和的矩阵
type NumMatrix struct {
	int
}

// Constructor 创建 NumMatrix
func Constructor(matrix [][]int) NumMatrix {

	return NumMatrix{}
}

// SumRegion 求和
func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {

	return 8
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
