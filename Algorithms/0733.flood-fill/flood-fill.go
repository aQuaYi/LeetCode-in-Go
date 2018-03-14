package problem0733

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func floodFill(image [][]int, sr, sc, newColor int) [][]int {
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}

	m, n := len(image), len(image[0])

	coordinates := make([][]int, 1, m*n)
	coordinates[0] = []int{sr, sc}

	for len(coordinates) > 0 {
		c := coordinates[0]
		coordinates = coordinates[1:]
		image[c[0]][c[1]] = newColor

		for i := 0; i < 4; i++ {
			x := c[0] + dx[i]
			y := c[1] + dy[i]
			if 0 <= x && x < m &&
				0 <= y && y < n &&
				image[x][y] == oldColor {
				coordinates = append(coordinates, []int{x, y})
			}
		}
	}

	return image
}
