package problem0913

// ref: https://leetcode.com/problems/cat-and-mouse/discuss/176177/Most-of-the-DFS-solutions-are-WRONG-check-this-case

func catMouseGame(graph [][]int) int {
	n := len(graph)
	color := [50][50][2]int{}
	outdegree := [50][50][2]int{}

	for i := 0; i < n; i++ { // cat
		for j := 0; j < n; j++ { // mouse
			outdegree[i][j][0] = len(graph[j])
			outdegree[i][j][1] = len(graph[i])
			for _, k := range graph[i] {
				if k == 0 {
					outdegree[i][j][1]--
					break
				}
			}
		}
	}

	q := make([][]int, 0, n)
	for k := 1; k < n; k++ {
		for m := 0; m < 2; m++ {
			color[k][0][m] = 1
			q = append(q, []int{k, 0, m, 1})
			color[k][k][m] = 2
			q = append(q, []int{k, k, m, 2})
		}
	}

	var cur []int

	for len(q) > 0 {
		cur, q = q[0], q[1:]
		cat, mouse, mouseMove, c := cur[0], cur[1], cur[2], cur[3]
		if cat == 2 &&
			mouse == 1 &&
			mouseMove == 0 {
			return c
		}
		prevMouseMove := 1 - mouseMove
		animal := mouse
		if prevMouseMove == 1 {
			animal = cat
		}
		for _, prev := range graph[animal] {
			prevCat := cat
			prevMouse := prev
			if prevMouseMove == 1 {
				prevCat = prev
				prevMouse = mouse
			}

			if prevCat == 0 {
				continue
			}

			if color[prevCat][prevMouse][prevMouseMove] > 0 {
				continue
			}

			outdegree[prevCat][prevMouse][prevMouseMove]--
			if (prevMouseMove == 1 && c == 2) ||
				(prevMouseMove == 0 && c == 1) ||
				(outdegree[prevCat][prevMouse][prevMouseMove] == 0) {
				color[prevCat][prevMouse][prevMouseMove] = c
				q = append(q, []int{prevCat, prevMouse, prevMouseMove, c})
			}
		}
	}

	return color[2][1][0]
}
