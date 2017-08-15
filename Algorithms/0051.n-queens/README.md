# [51. N-Queens](https://leetcode.com/problems/n-queens/)

## 题目
The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.

![8-queens](8-queens.png)

Given an integer n, return all distinct solutions to the n-queens puzzle.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.

For example,
There exist two distinct solutions to the 4-queens puzzle:
```
[
 [".Q..",  // Solution 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // Solution 2
  "Q...",
  "...Q",
  ".Q.."]
]
```
## 解题思路
题目要求，在以下地方都没有其他的Queen
1. 所在的行
1. 所在的列
1. 两条45度对角线上

和前面的数独问题一样的解题思路
## 总结


