package problem0923

// refer: https://leetcode.com/problems/3sum-with-multiplicity/discuss/181131/C++JavaPython-O(1012)

const mod = 1e9 + 7

// first of all, swap will NOT change answer
// for example
//     A[0]=0, A[1]=1, A[2]=2 , A[0]+A[1]+A[2]=3
//     A[0]=1, A[1]=0, A[2]=2 , A[0]+A[1]+A[2]=3
// i, j, k and A[i]+A[j]+A[k] don't change, answer don't change
// if we sort A
// then, i<j<k => A[i]<=A[j]<=A[k]
// but we do not need sort A
// just count it

func threeSumMulti(A []int, target int) int {
	count := [101]int{}
	for _, a := range A {
		count[a]++
	}

	res := 0

	for Ai := 0; Ai <= 100; Ai++ {
		for Aj := Ai; Aj <= 100; Aj++ { // so always Ai<=Aj
			Ak := target - Ai - Aj
			if Ak < 0 || 100 < Ak { // need 0<= A[k] <= 100
				continue
			}

			// remember, i<j<k => A[i]<=A[j]<=A[k]
			// A[i]<=A[j]<=A[k] has 4 conditions
			// 按照以下情况分别计算组合数的话
			// 可以确保 A[i]<=A[j]<=A[k] => i<j<k
			switch {
			case Ai == Aj && Aj == Ak:
				// combination: count[Ai] choose 3
				res += count[Ai] * (count[Ai] - 1) * (count[Ai] - 2) / 6
			case Ai == Aj && Aj < Ak:
				// ( count[Ai] choose 2 ) multiply count[Ak]
				res += count[Ai] * (count[Ai] - 1) / 2 * count[Ak]
			case Aj == Ak: // Ai < Aj = Ak
				// count[Ai] multiply ( count[Aj] choose 2 )
				res += count[Ai] * count[Aj] * (count[Aj] - 1) / 2
			case Aj < Ak: // Ai < Aj < Ak
				res += count[Ai] * count[Aj] * count[Ak]
			}
		}
	}

	return res % mod
}
