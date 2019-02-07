package problem0952

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313}

func largestComponentSize(A []int) int {
	size := len(A)
	u := newUnion(size)

	// // TODO: 删除此处内容
	// sort.Ints(A)

	// rec = map[factor]index
	rec := make(map[int]int, size*2)

	for i := 0; i < size; i++ {
		a := A[i]
		// for _, p := range primes {
		// if p*p > a {
		// break
		// }
		for p := 2; p*p <= a; p++ {
			if a%p != 0 {
				continue
			}
			if j, ok := rec[p]; ok {
				u.union(i, j)
			} else {
				rec[p] = i
			}
			d := a / p
			if j, ok := rec[d]; ok {
				u.union(i, j)
			} else {
				rec[d] = i
			}
		}
		if j, ok := rec[a]; ok {
			u.union(i, j)
		} else {
			rec[a] = i
		}
	}

	// fmt.Println("===")
	// for k, v := range rec {
	// 	if k == 29 {
	// 		fmt.Println(k, v)
	// }
	// }

	// for i := 0; i < size; i++ {
	// 	star := ""
	// 	if u.parents[i] != 1 {
	// 		star = "*"
	// 	}
	// 	ej := ""
	// 	if A[i]%29 == 0 {
	// 		ej = "$"
	// 	}
	// 	fmt.Println(i, u.parents[i], A[i], star, ej)
	// }

	return u.max
}

// Robert Sedgewick 算法（第4版） 1.5.2.7
// union-find (加权 quick-union)，还作了路径压缩优化

type union struct {
	parents []int // 父链接数组(由触点索引)
	sizes   []int // (由触点索引的) 各个根节点所对应的分量的大小
	max     int
}

func newUnion(N int) *union {
	parents := make([]int, N)
	for i := range parents {
		parents[i] = i
	}
	counts := make([]int, N)
	for i := range counts {
		counts[i] = 1
	}
	return &union{
		parents: parents,
		sizes:   counts,
		max:     1,
	}
}

func (u *union) isConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *union) find(p int) int {
	// 跟随连接找到根节点
	if u.parents[p] != p {
		u.parents[p] = u.find(u.parents[p])
	}
	return u.parents[p]
}

func (u *union) union(p, q int) {
	i, j := u.find(p), u.find(q)
	if i == j {
		return
	}

	if u.sizes[i] > u.sizes[j] {
		i, j = j, i
	}

	// 将小树的根节点连接到大树的根节点
	u.parents[i] = j
	u.sizes[j] += u.sizes[i]
	u.max = max(u.max, u.sizes[j])
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
