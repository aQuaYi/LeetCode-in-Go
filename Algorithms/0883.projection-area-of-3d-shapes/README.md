# [883. Projection Area of 3D Shapes](https://leetcode.com/problems/projection-area-of-3d-shapes/)

## 题目

On aN*N grid, we place some1 * 1 * 1cubes that are axis-aligned with the x, y, and z axes.

Each valuev = grid[i][j]represents a tower ofvcubes placed on top of grid cell (i, j).

Now we view theprojectionof these cubesonto the xy, yz, and zx planes.

A projection is like a shadow, thatmaps our 3 dimensional figure to a 2 dimensional plane.

Here, we are viewing the "shadow" when looking at the cubes from the top, the front, and the side.

Return the total area of all three projections.

Example 1:

```text
Input: [[2]]
Output: 5
```

Example 2:

```text
Input: [[1,2],[3,4]]
Output: 17
Explanation:
Here are the three projections ("shadows") of the shape made with each axis-aligned plane.
```

![shadow](shadow.png)

Example 3:

```text
Input: [[1,0],[0,2]]
Output: 8
```

Example 4:

```text
Input: [[1,1,1],[1,0,1],[1,1,1]]
Output: 14
```

Example 5:

```text
Input: [[2,2,2],[2,1,2],[2,2,2]]
Output: 21
```

Note:

1. 1 <= grid.length = grid[0].length<= 50
1. 0 <= grid[i][j] <= 50

## 解题思路

见程序注释
