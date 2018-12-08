package problem0913

// ref: https://leetcode.com/problems/cat-and-mouse/discuss/176177/Most-of-the-DFS-solutions-are-WRONG-check-this-case

func catMouseGame(graph [][]int) int {
	n := len(graph)
	colors := [50][50][2]int{}
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

	queue := make([][4]int, 0, n)

	colorizeAndPush := func(cat, mouse, mouseMove, color int) {
		colors[cat][mouse][mouseMove] = color
		queue = append(queue, [4]int{cat, mouse, mouseMove, color})
	}

	for k := 1; k < n; k++ {
		for m := 0; m < 2; m++ {
			colorizeAndPush(k, 0, m, 1)
			colorizeAndPush(k, k, m, 2)
		}
	}

	var cur [4]int

	for len(queue) > 0 {
		cur, queue = queue[0], queue[1:]
		cat, mouse, mouseMove, color := cur[0], cur[1], cur[2], cur[3]
		if cat == 2 &&
			mouse == 1 &&
			mouseMove == 0 {
			return color
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

			if colors[prevCat][prevMouse][prevMouseMove] > 0 {
				continue
			}

			outdegree[prevCat][prevMouse][prevMouseMove]--
			if (prevMouseMove == 1 && color == 2) ||
				(prevMouseMove == 0 && color == 1) ||
				(outdegree[prevCat][prevMouse][prevMouseMove] == 0) {
				colors[prevCat][prevMouse][prevMouseMove] = color
				queue = append(queue, [4]int{prevCat, prevMouse, prevMouseMove, color})
			}
		}
	}

	return colors[2][1][0]
}
