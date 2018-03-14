package problem0357

var res = []int{1, 10, 91, 739, 5275, 32491, 168571, 712891, 2345851, 5611771, 8877691}

func countNumbersWithUniqueDigits(n int) int {
	if n >= 10 {
		return res[10]
	}
	return res[n]
}

// 参考 《代码大全2》 第 18 章 表驱动法
