# [832. Flipping an Image](https://leetcode.com/problems/flipping-an-image/)

## 题目

Given a binary matrix A, we want to flip the image horizontally, then invert it, and return the resulting image.

To flip an image horizontally means that each row of the image is reversed. For example, flipping[1, 1, 0]horizontally results in[0, 1, 1].

To invert an image meansthat each 0 is replaced by 1, and each 1 is replaced by 0.For example, inverting[0, 1, 1]results in[1, 0, 0].

Example 1:

```text
Input: [[1,1,0],[1,0,1],[0,0,0]]
Output: [[1,0,0],[0,1,0],[1,1,1]]
Explanation: First reverse each row: [[0,1,1],[1,0,1],[0,0,0]].
Then, invert the image: [[1,0,0],[0,1,0],[1,1,1]]
```

Example 2:

```text
Input: [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
Output: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
Explanation: First reverse each row: [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]].
Then invert the image: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
```

Notes:

1. 1 <= A.length = A[0].length <= 20
1. 0 <= A[i][j]<=1

## 解题思路

见程序注释
