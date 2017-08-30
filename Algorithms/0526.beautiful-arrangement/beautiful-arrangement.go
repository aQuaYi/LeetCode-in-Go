package Problem0526

func countArrangement(N int) int {
	rec := make([]bool, N)
	res := 0

	var dfs func(int)
	dfs = func(level int) {
		if level == N+1 {
			res++
			return
		}

		for i := range rec {
			if rec[i] == false && 
			((i+1)%level == 0 ||level%(i+1)==0) {
				rec[i] = true
				dfs(level + 1)
				rec[i] = false
			}
		}
	}

	dfs(1)
	return res
}
