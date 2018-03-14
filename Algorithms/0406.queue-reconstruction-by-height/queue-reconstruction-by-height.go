package problem0406

import "sort"

func reconstructQueue(people [][]int) [][]int {
	res := make([][]int, 0, len(people))

	// 按照 persons 的排序方式，对 people 进行排序
	sort.Sort(persons(people))

	// 把 person 插入到 res[idx] 上
	insert := func(idx int, person []int) {
		res = append(res, person)
		// 插入到末尾
		if len(res)-1 == idx {
			return
		}
		// 插入到中间位置
		copy(res[idx+1:], res[idx:])
		res[idx] = person
	}

	// 对于 res[i] 来说，
	// 如果把 res[:i] 中所有 h < res[i] 的元素全部删除后，得到的 res'
	// len(res') == res[i][1]
	//
	//                   0      1      2      3      4      5
	// [4,k] 的 k 在 [[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]] 中与其索引号一致
	// [5,k] 的 k 在 [[5,0], [7,0], [5,2], [6,1], [7,1]]        中与其索引号一致
	// [6,k] 的 k 在 [[7,0], [6,1], [7,1]]                      中与其索引号一致
	// [7,k] 的 k 在 [[7,0], [7,1]]                             中与其索引号一致
	// 下面的 for 循环，就是上面的删除的逆过程
	for _, p := range people {
		insert(p[1], p)
	}

	return res
}

// persons 实现了 sort.Interface 接口
type persons [][]int

func (p persons) Len() int { return len(p) }

// 以 h 的降序为主
// 以 k 的升序为辅
func (p persons) Less(i, j int) bool {
	if p[i][0] == p[j][0] {
		return p[i][1] < p[j][1]
	}
	return p[i][0] > p[j][0]
}

func (p persons) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
