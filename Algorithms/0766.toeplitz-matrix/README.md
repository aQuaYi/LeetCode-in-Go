# [766. Toeplitz Matrix](https://leetcode.com/problems/toeplitz-matrix/)

## 题目

A matrix is Toeplitz if every diagonal from top-left to bottom-right has the same element.

Now given an `M x N` matrix, return `True` if and only if the matrix is Toeplitz.

Example 1:

```text
Input: matrix = [[1,2,3,4],[5,1,2,3],[9,5,1,2]]
Output: True
Explanation:
1234
5123
9512

In the above grid, the&nbsp;diagonals are &quot;[9]&quot;, &quot;[5, 5]&quot;, &quot;[1, 1, 1]&quot;, &quot;[2, 2, 2]&quot;, &quot;[3, 3]&quot;, &quot;[4]&quot;, and in each diagonal all elements are the same, so the answer is True.
```

Example 2:

```text
Input: matrix = [[1,2],[2,2]]
Output: False
Explanation:
The diagonal &quot;[1, 2]&quot; has different elements.
```

Note:

1. `matrix` will be a 2D array of integers.
1. `matrix` will have a number of rows and columns in range [1, 20].
1. `matrix[i][j]` will be integers in range [0, 99].

## 解题思路

见程序注释
