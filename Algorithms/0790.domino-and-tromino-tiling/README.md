# [790. Domino and Tromino Tiling](https://leetcode.com/problems/domino-and-tromino-tiling/)

## 题目

We have two types of tiles: a 2x1 domino shape, and an "L" tromino shape. These shapes may be rotated.

```text
XX  <- domino

XX  <- "L" tromino
X
```

Given N, how many ways are there to tile a 2 x N board? Return your answer modulo 10^9 + 7.

(In a tiling, every square must be covered by a tile. Two tilings are different if and only if there are two 4-directionally adjacent cells on the board such that exactly one of the tilings has both squares occupied by a tile.)

```text
Example:
Input: 3
Output: 5
Explanation:
The five different ways are listed below, different letters indicates different tiles:
XYZ XXZ XYY XXY XYY
XYZ YYZ XZZ XYY XXY
```

Note:

1. N will be in range [1, 1000].

## 解题思路

见程序注释
