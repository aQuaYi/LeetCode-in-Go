# [959. Regions Cut By Slashes](https://leetcode.com/problems/regions-cut-by-slashes/)

In a N x N grid composed of 1 x 1 squares, each 1 x 1 square consists of a /, \, or blank space.  These characters divide the square into contiguous regions.

(Note that backslash characters are escaped, so a \ is represented as "\\".)

Return the number of regions.

Example 1:

```text
Input:
[
  " /",
  "/ "
]

Output: 2
Explanation: The 2x2 grid is as follows:
```

![1](1.png)

Example 2:

```text
Input:
[
  " /",
  "  "
]
Output: 1
Explanation: The 2x2 grid is as follows:
```

![2](2.png)

Example 3:

```text
Input:
[
  "\\/",
  "/\\"
]
Output: 4
Explanation: (Recall that because \ characters are escaped, "\\/" refers to \/, and "/\\" refers to /\.)
The 2x2 grid is as follows:
```

![3](3.png)

Example 4:

```text
Input:
[
  "/\\",
  "\\/"
]
Output: 5
Explanation: (Recall that because \ characters are escaped, "/\\" refers to /\, and "\\/" refers to \/.)
The 2x2 grid is as follows:
```

![4](4.png)

Example 5:

```text
Input:
[
  "//",
  "/ "
]
Output: 3
Explanation: The 2x2 grid is as follows:
```

![5](5.png)

Note:

1. 1 <= grid.length == grid[0].length <= 30
1. grid[i][j] is either '/', '\', or ' '.