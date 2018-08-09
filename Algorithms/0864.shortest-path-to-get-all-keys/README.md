# [864. Shortest Path to Get All Keys](https://leetcode.com/problems/shortest-path-to-get-all-keys/)

## 题目

We are given a 2-dimensionalgrid."." is an empty cell, "#" isa wall, "@" is the starting point, ("a", "b", ...) are keys, and ("A","B", ...) are locks.

We start at the starting point, and one move consists of walking one space in one of the 4 cardinal directions. We cannot walk outside the grid, or walk into a wall. If we walk over a key, we pick it up. We can't walk over a lock unless we have the corresponding key.

For some 1 <= K <= 6, there is exactly one lowercase and one uppercase letter of the first K letters of the English alphabet in the grid. This means that there is exactly one key for each lock, and one lock for each key; and also that the letters used to represent the keys and locks werechosen in the same order as the English alphabet.

Return the lowest number of moves to acquire all keys. Ifit's impossible, return -1.

Example 1:

```text
Input: ["@.a.#","###.#","b.A.B"]
Output: 8
```

Example 2:

```text
Input: ["@..aA","..B#.","....b"]
Output: 6
```

Note:

1. 1 <= grid.length<= 30
1. 1 <= grid[0].length<= 30
1. grid[i][j] contains only '.', '#', '@','a'-'f' and 'A'-'F'
1. The number of keys is in [1, 6]. Each key has a different letter and opens exactly one lock.

## 解题思路

见程序注释
